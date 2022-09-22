package domain

import (
	"context"
	"math/big"
	"time"
)

// Staff entity
type Staff struct {
	ID          big.Int    `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	StaffCode   string     `json:"staff_code"`
	BirthDate   string     `json:"birth_date"`
	Gender      int        `json:"gender"`
	PhoneNumber string     `json:"phone_number"`
	UnitId      int        `json:"unit_id"`
	StaffTypeId int        `json:"staff_type_id"`
	Address     string     `json:"address"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// StaffQueryParam for search
type StaffQueryParam struct {
	Email string `json:"email"`
}

// StaffRepository represent the staff's usecases
type StaffRepository interface {
	Fetch(context.Context) ([]Staff, error)
	Find(ctx context.Context, id int) (*Staff, error)
	Store(ctx context.Context, u *Staff) error
	Search(ctx context.Context, q StaffQueryParam) ([]Staff, error)
	FindByQuery(ctx context.Context, q StaffQueryParam) (*Staff, error)
}

// StaffUseCase represent the staff's repository contract
type StaffUsecase interface {
	Fetch(context.Context) ([]Staff, error)
	Find(ctx context.Context, id int) (*Staff, error)
	Store(ctx context.Context, u *Staff) error
	Search(ctx context.Context, q StaffQueryParam) ([]Staff, error)
	FindByQuery(ctx context.Context, q StaffQueryParam) (*Staff, error)
}
