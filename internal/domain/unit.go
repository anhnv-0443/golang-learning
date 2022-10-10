package domain

import (
	"context"
	"math/big"
)

// Unit entity
type Unit struct {
	ID   big.Int `json:"id"`
	Name string  `json:"name"`
}

// UnitRepository represent the Unit's usecases
type UnitRepository interface {
	Fetch(context.Context) ([]Unit, error)
	Find(ctx context.Context, id int) (*Unit, error)
	Store(ctx context.Context, u *Unit) error
}

// UnitUseCase represent the Unit's repository contract
type UnitUseCase interface {
	Fetch(context.Context) ([]Unit, error)
	Find(ctx context.Context, id int) (*Unit, error)
	Store(ctx context.Context, u *Unit) error
}
