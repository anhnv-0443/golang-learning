package http

import (
	"go-app/internal/domain"
	"net/http"

	"fmt"

	"github.com/labstack/echo/v4"
)

// GithubHandler represent the http handler
type GithubHandler struct {
	Usecase domain.GithubUsecase
}

func (hl *GithubHandler) ConnectGithub(c echo.Context) error {
	userReq := new(GithubConnectRequest)

	if err := c.Bind(userReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	err := hl.Usecase.RequestCodeToGithub(ctx, userReq.ClientId, "")
	if err != nil {
		return c.JSON(http.StatusOK, &GithubResponse{false})
	}

	return c.JSON(http.StatusOK, &GithubResponse{true})
}

func (hl *GithubHandler) CallbackFromGithub(c echo.Context) error {
	userReq := new(GithubCallbackRequest)

	if err := c.Bind(userReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	fmt.Println(userReq.Code)
	return c.JSON(http.StatusOK, &GithubResponse{true})
}

func NewHandler(g *echo.Group, uc domain.GithubUsecase) {
	handler := &GithubHandler{
		Usecase: uc,
	}

	g.GET("/connect", handler.ConnectGithub)
	g.GET("/callback", handler.CallbackFromGithub)
}
