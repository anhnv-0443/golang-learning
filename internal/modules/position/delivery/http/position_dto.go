package http

import (
	"go-app/internal/domain"
	"time"
)

// PositionsResponse is array of position response
type PositionsResponse []PositionResponse

// PositionRequest is struct used for create st
type PositionRequest struct {
	Name string
}

// PositionResponse is struct used for position
type PositionResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ErrorResponse is struct when error
type ErrorResponse struct {
	Message string `json:"message"`
}

// StatusResponse is struct when success
type StatusResponse struct {
	Status bool `json:"status"`
	Data   *PositionRequest
}

// ConvertPositionToResponse DTO
func ConvertPositionToResponse(positions *domain.Position) PositionResponse {
	return PositionResponse{
		ID:        int(positions.ID),
		Name:      positions.Name,
		CreatedAt: positions.CreatedAt,
		UpdatedAt: positions.UpdatedAt,
	}
}

//// ConvertPositionsToResponse DTO
func ConvertPositionsToResponse(Positions []domain.Position) PositionsResponse {
	positionsRes := make(PositionsResponse, 0)

	for _, u := range Positions {
		positionRes := PositionResponse{
			ID:        int(u.ID),
			Name:      u.Name,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}

		positionsRes = append(positionsRes, positionRes)
	}

	return positionsRes
}

//
//// ConvertRequestToEntity DTO
func ConvertRequestToEntity(u *PositionRequest) *domain.Position {
	return &domain.Position{
		Name: u.Name,
	}
}
