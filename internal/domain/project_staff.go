package domain

import (
	"context"
	"math/big"
)

// ProjectStaff entity
type ProjectStaff struct {
	StaffId big.Int `json:"id"`
}

// ProjectStaffRepository represent the ProjectStaff's usecases
type ProjectStaffRepository interface {
	Fetch(context.Context) ([]ProjectStaff, error)
	Find(ctx context.Context, id int) (*ProjectStaff, error)
	Store(ctx context.Context, u *ProjectStaff) error
}

// ProjectStaffUseCase represent the ProjectStaff's repository contract
type ProjectStaffUseCase interface {
	Fetch(context.Context) ([]ProjectStaff, error)
	Find(ctx context.Context, id int) (*ProjectStaff, error)
	Store(ctx context.Context, u *ProjectStaff) error
}
