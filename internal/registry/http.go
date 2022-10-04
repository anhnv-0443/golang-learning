package registry

import (
	"go-app/config"
	"go-app/internal/domain"
	authHttp "go-app/internal/modules/auth/delivery/http"
	roleHttp "go-app/internal/modules/role/delivery/http"
	slackHttp "go-app/internal/modules/slack/delivery/http"
	staffHttp "go-app/internal/modules/staff/delivery/http"
	userHttp "go-app/internal/modules/user/delivery/http"
	"go-app/pkg/validate"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewHTTPHandler registry http
func NewHTTPHandler(e *echo.Echo, uc *Usecase) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = validate.NewValidate()

	apiGroup := e.Group("/api")
	slackGroup := e.Group("/slack")

	authGroup := apiGroup.Group("")

	authHttp.NewHandler(authGroup, uc.AuthUsecase)

	DefaultJWTConfig := middleware.JWTConfig{
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
		Claims:      &domain.Claims{},
		SigningKey:  []byte(config.GetAppConfig().AppJWTKey),
	}
	apiGroup.Use(middleware.JWTWithConfig(DefaultJWTConfig))

	// CORS restricted with a custom function to allow origins
	// and with the GET, PUT, POST or DELETE methods allowed.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: corsAllowOrigin,
		AllowMethods:    []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	staffHttp.NewHandler(apiGroup, uc.StaffUsecase)
	// Middleware verifying request from slack
	slackGroup.Use(slackHttp.SlackVerifyingRequests(config.GetSlackConfig()))

	roleHttp.NewHandler(apiGroup, uc.RoleUsecase)
	userHttp.NewHandler(apiGroup, uc.UserUsecase)

	slackHttp.NewHandler(slackGroup, uc.SlackUsecase)
}

func corsAllowOrigin(origin string) (bool, error) {
	list := strings.Split(config.GetAppConfig().AllowedOrigin, ",")

	for _, b := range list {
		if b == origin {
			return true, nil
		}
	}

	return false, nil
}
