package domain

import (
	"context"
	"math/big"
)

// StaffType entity
type StaffType struct {
	ID   big.Int `json:"id"`
	Name string  `json:"name"`
}

// StaffTypeRepository represent the StaffType's usecases
type StaffTypeRepository interface {
	Fetch(context.Context) ([]StaffType, error)
	Find(ctx context.Context, id int) (*StaffType, error)
	Store(ctx context.Context, u *StaffType) error
}

// StaffTypeUseCase represent the StaffType's repository contract
type StaffTypeUseCase interface {
	Fetch(context.Context) ([]StaffType, error)
	Find(ctx context.Context, id int) (*StaffType, error)
	Store(ctx context.Context, u *StaffType) error
}
