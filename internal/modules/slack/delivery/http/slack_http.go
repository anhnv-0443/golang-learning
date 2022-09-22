package http

import (
	"go-app/internal/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

// SlackHandler represent the httphandler
type SlackHandler struct {
	Usecase domain.SlackUsecase
}

// NewHandler will initialize the Slack endpoint
func NewHandler(g *echo.Group, uc domain.SlackUsecase) {
	handler := &SlackHandler{
		Usecase: uc,
	}

	g.POST("/", handler.HandleSlackRequest)
}

// HandleSlackRequest handle request from slack
func (hl *SlackHandler) HandleSlackRequest(c echo.Context) error {
	userReq := new(SlackRequest)
	if err := c.Bind(userReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()

	switch {
	case userReq.Type == "url_verification":
		return c.JSON(http.StatusOK, &SlackChallengeResponse{userReq.Challenge})
	case userReq.Type == "event_callback" && userReq.Event.Type == "app_mention":
		err := hl.Usecase.HandleAppMentionEvent(ctx, ConvertRequestToSlackAppMentionEvent(&userReq.Event))
		if err != nil {
			return c.JSON(http.StatusOK, &SlackResponse{false})
		}
	}

	return c.JSON(http.StatusOK, &SlackResponse{true})
}
