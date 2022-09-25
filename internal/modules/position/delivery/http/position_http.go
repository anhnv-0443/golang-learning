package http

import (
	"github.com/labstack/echo/v4"
	"go-app/internal/domain"
	"net/http"
)

// UserHandler represent the httphandler
type PositionHandler struct {
	Usecase domain.PositionUsecase
}

// NewHandler will initialize the Positions/ resources endpoint
func NewHandler(g *echo.Group, uc domain.PositionUsecase) {
	handler := &PositionHandler{
		Usecase: uc,
	}

	g.GET("/positions", handler.Index)
	g.POST("/positions", handler.Store)
}

// Index will fetch data
func (hl *PositionHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	positions, _ := hl.Usecase.Fetch(ctx)

	return c.JSON(http.StatusOK, ConvertPositionsToResponse(positions))
}

// Store will create data
func (hl *PositionHandler) Store(c echo.Context) error {
	positionReq := new(PositionRequest)

	if err := c.Bind(positionReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	err := c.Validate(positionReq)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	position := ConvertRequestToEntity(positionReq)

	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, position); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, StatusResponse{Status: true, Data: positionReq})
}
