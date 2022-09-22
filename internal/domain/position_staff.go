package domain

import (
	"context"
	"math/big"
)

// PositionStaff entity
type PositionStaff struct {
	ProjectId big.Int `json:"project_id"`
	StaffId   big.Int `json:"staff_id"`
}

// PositionStaffRepository represent the PositionStaff's usecases
type PositionStaffRepository interface {
	Fetch(context.Context) ([]PositionStaff, error)
	Find(ctx context.Context, id int) (*PositionStaff, error)
	Store(ctx context.Context, u *PositionStaff) error
}

// PositionStaffUseCase represent the PositionStaff's repository contract
type PositionStaffUseCase interface {
	Fetch(context.Context) ([]PositionStaff, error)
	Find(ctx context.Context, id int) (*PositionStaff, error)
	Store(ctx context.Context, u *PositionStaff) error
}
