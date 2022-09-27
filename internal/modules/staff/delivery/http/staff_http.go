package http

import (
	"go-app/internal/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserHandler represent the httphandler
type StaffHandler struct {
	Usecase                  domain.StaffUsecase
	StaffLanguageCodeUsecase domain.StaffLanguageCodeUseCase
}

// NewHandler will initialize the Staffs/ resources endpoint
func NewHandler(g *echo.Group, uc domain.StaffUsecase) {
	handler := &StaffHandler{
		Usecase: uc,
	}

	g.GET("/staffs", handler.Index)
	g.GET("/staffs/:id", handler.Show)
	g.POST("/staffs", handler.Store)
	g.PATCH("/staffs/:id", handler.Update)
	g.DELETE("/staffs/:id", handler.Delete)
}

// Index will fetch data
func (hl *StaffHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	staffs, _ := hl.Usecase.Fetch(ctx)

	return c.JSON(http.StatusOK, ConvertStaffsToResponse(staffs))
}

// Show will Find data
func (hl *StaffHandler) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Message: err.Error()})
	}
	ctx := c.Request().Context()
	staff, err := hl.Usecase.Find(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ConvertStaffToResponse(staff))
}

// Store will create data
func (hl *StaffHandler) Store(c echo.Context) error {
	staffReq := new(StaffRequest)

	if err := c.Bind(staffReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	err := c.Validate(staffReq)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	staff := ConvertRequestToEntity(staffReq)

	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, staff); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, StatusResponse{Status: true})
}

// Update will update data
func (hl *StaffHandler) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}

// Delete will delete data
func (hl *StaffHandler) Delete(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
