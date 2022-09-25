package http

import (
	"github.com/labstack/echo/v4"
	"go-app/internal/domain"
	"net/http"
)

// UserHandler represent the httphandler
type LanguageCodeHandler struct {
	Usecase domain.LanguageCodeUsecase
}

// NewHandler will initialize the LanguageCodes/ resources endpoint
func NewHandler(g *echo.Group, uc domain.LanguageCodeUsecase) {
	handler := &LanguageCodeHandler{
		Usecase: uc,
	}

	g.GET("/language_code", handler.Index)
	g.POST("/language_code", handler.Store)
}

// Index will fetch data
func (hl *LanguageCodeHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	language_codes, _ := hl.Usecase.Fetch(ctx)

	return c.JSON(http.StatusOK, ConvertLanguageCodesToResponse(language_codes))
}

func (hl *LanguageCodeHandler) Store(c echo.Context) error {
	languageCodeReq := new(LanguageCodeRequest)

	if err := c.Bind(languageCodeReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	err := c.Validate(languageCodeReq)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	languageCode := ConvertRequestToEntity(languageCodeReq)

	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, languageCode); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, StatusResponse{Status: true, Data: languageCodeReq})
}
