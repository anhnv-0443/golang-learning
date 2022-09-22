package domain

import (
	"context"
	"math/big"
)

// Position entity
type Position struct {
	ID   big.Int `json:"id"`
	Name string  `json:"name"`
}

// PositionRepository represent the Position's usecases
type PositionRepository interface {
	Fetch(context.Context) ([]Position, error)
	Find(ctx context.Context, id int) (*Position, error)
	Store(ctx context.Context, u *Position) error
}

// PositionUseCase represent the Position's repository contract
type PositionUseCase interface {
	Fetch(context.Context) ([]Position, error)
	Find(ctx context.Context, id int) (*Position, error)
	Store(ctx context.Context, u *Position) error
}
