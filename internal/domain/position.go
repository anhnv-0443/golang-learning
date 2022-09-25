package domain

import (
	"context"
	"time"
)

// Position entity
type Position struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// PositionRepository represent the Position's usecases
type PositionRepository interface {
	Fetch(context.Context) ([]Position, error)
	Find(ctx context.Context, id int) (*Position, error)
	Store(ctx context.Context, u *Position) error
}

// PositionUseCase represent the Position's repository contract
type PositionUsecase interface {
	Fetch(context.Context) ([]Position, error)
	Find(ctx context.Context, id int) (*Position, error)
	Store(ctx context.Context, u *Position) error
}
