package http_test

import (
	"go-app/config"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	slackHttp "go-app/internal/modules/slack/delivery/http"
)

type slackVerifyingRequestsTestCase struct {
	name             string
	requestTimestamp string
	requestSignature string
	requestBody      string
	expectedCode     int
	expectedResBody  string
	expectedError    error
}

func TestSlackVerifyingRequests(t *testing.T) {
	e := echo.New()
	t.Parallel()
	sampleRequestBody := `token=xyzz0WbapA4vBCDEFasx0q6G&team_id=T1DC2JH3J&team_domain=testteamnow&channel_id=G8PSS9T3V` +
		`&channel_name=foobar&user_id=U2CERLKJA&user_name=roadrunner&command=%2Fwebhook-collect&text=&` +
		`response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT1DC2JH3J%2F397700885554%2F96rGlfmibIGlgcZRskXaIFfN&` +
		`trigger_id=398738663015.47445629121.803a0bc887a14d10d2c447fce8b6703c`

	tests := []slackVerifyingRequestsTestCase{
		{
			name:             "Invalid request timestamp",
			requestTimestamp: "1131420618",
			requestSignature: "v0=a2114d57b48eac39b9ad189dd8316235a7b4a8d21a10bd27519666489c69b503",
			requestBody:      sampleRequestBody,
			expectedError:    echo.NewHTTPError(http.StatusBadRequest, "request timestamp invalid"),
		},
		{
			name:             "Invalid request timestamp",
			requestTimestamp: "1831420618",
			requestSignature: "v0=a2114d57b48eac39b9ad189dd8316235a7b4a8d21a10bd27519666489c69b503",
			requestBody:      sampleRequestBody,
			expectedError:    echo.NewHTTPError(http.StatusBadRequest, "request timestamp invalid"),
		},
		{
			name:             "Invalid request signature",
			requestTimestamp: "1531420618",
			requestSignature: "v0=a2114d57b48eac39b9ad189dd8316235a7b4a8d219b503",
			requestBody:      sampleRequestBody,
			expectedError:    echo.NewHTTPError(http.StatusBadRequest, "request signature invalid"),
		},
		{
			name:             "Valid Request",
			requestTimestamp: "1531420618",
			requestSignature: "v0=a2114d57b48eac39b9ad189dd8316235a7b4a8d21a10bd27519666489c69b503",
			requestBody:      sampleRequestBody,
			expectedCode:     http.StatusOK,
			expectedResBody:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub time.Now() function
			slackHttp.TimeNow = func() time.Time {
				return time.Unix(1531420618, 0)
			}

			req := httptest.NewRequest(echo.POST, "/slack/", strings.NewReader(tc.requestBody))
			req.Header.Set("X-Slack-Request-Timestamp", tc.requestTimestamp)
			req.Header.Set("X-Slack-Signature", tc.requestSignature)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			middleware := slackHttp.SlackVerifyingRequests(config.SlackConfig{SigningSecret: "8f742231b10e8888abcd99yyyzzz85a5"})
			handler := middleware(echo.HandlerFunc(func(c echo.Context) error {
				return c.NoContent(200)
			}))

			err := handler(c)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.Equal(t, tc.expectedResBody, rec.Body.String())
				assert.Equal(t, tc.expectedCode, rec.Code)

				// Should keep request body after read
				bodyBytes, _ := io.ReadAll(c.Request().Body)
				reqBody := string(bodyBytes)
				assert.Equal(t, tc.requestBody, reqBody)
			}
		})
	}
}
