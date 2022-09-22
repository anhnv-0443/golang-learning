package domain

import (
	"context"
	"math/big"
)

// ProjectPosition entity
type ProjectPosition struct {
	ProjectId  big.Int `json:"project_id"`
	PositionId string  `json:"position_id"`
}

// ProjectPositionRepository represent the ProjectPosition's usecases
type ProjectPositionRepository interface {
	Fetch(context.Context) ([]ProjectPosition, error)
	Find(ctx context.Context, id int) (*ProjectPosition, error)
	Store(ctx context.Context, u *ProjectPosition) error
}

// ProjectPositionUseCase represent the ProjectPosition's repository contract
type ProjectPositionUseCase interface {
	Fetch(context.Context) ([]ProjectPosition, error)
	Find(ctx context.Context, id int) (*ProjectPosition, error)
	Store(ctx context.Context, u *ProjectPosition) error
}
