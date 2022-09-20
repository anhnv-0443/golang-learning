package http

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"go-app/config"
	"go-app/internal/constants"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// TimeNow have to export this var for testing purpose
var TimeNow = time.Now

// SlackVerifyingRequests is Middleware verifying request from slack
func SlackVerifyingRequests(config config.SlackConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			timestamp := c.Request().Header.Get("X-Slack-Request-Timestamp")

			if !isValidTimestamp(timestamp) {
				return echo.NewHTTPError(http.StatusBadRequest, "request timestamp invalid")
			}

			bodyBytes, _ := io.ReadAll(c.Request().Body)
			reqBody := string(bodyBytes)

			c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			messageSignature := c.Request().Header.Get("X-Slack-Signature")
			expectedSignature := generateSignature(reqBody, timestamp, config.SigningSecret)

			if isValidSignature(messageSignature, expectedSignature) {
				return next(c)
			}

			return echo.NewHTTPError(http.StatusBadRequest, "request signature invalid")
		}
	}
}

// The request timestamp is more than five minutes from local time.
// It could be a replay attack, so let's ignore it.
func isValidTimestamp(timestamp string) bool {
	t, _ := strconv.ParseInt(timestamp, constants.DecimalNum, constants.Int64Num)
	currentTimestamp := TimeNow().Unix()
	fiveMinute := 300
	timeDiff := int(math.Abs(float64(currentTimestamp - t)))

	return timeDiff <= fiveMinute
}

func generateSignature(body string, timestamp string, secret string) string {
	sigBaseString := "v0:" + timestamp + ":" + body
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(sigBaseString))

	return "v0=" + hex.EncodeToString(mac.Sum(nil))
}

func isValidSignature(messageSignature, expectedSignature string) bool {
	return hmac.Equal([]byte(messageSignature), []byte(expectedSignature))
}
