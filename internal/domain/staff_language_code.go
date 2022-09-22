package domain

import (
	"context"
	"math/big"
)

// StaffLanguageCode entity
type StaffLanguageCode struct {
	StaffId        big.Int `json:"staff_id"`
	LanguageCodeId big.Int `json:"language_code_id"`
}

// StaffLanguageCodeRepository represent the StaffLanguageCode's usecases
type StaffLanguageCodeRepository interface {
	Fetch(context.Context) ([]StaffLanguageCode, error)
	Find(ctx context.Context, id int) (*StaffLanguageCode, error)
	Store(ctx context.Context, u *StaffLanguageCode) error
}

// StaffLanguageCodeUseCase represent the StaffLanguageCode's repository contract
type StaffLanguageCodeUseCase interface {
	Fetch(context.Context) ([]StaffLanguageCode, error)
	Find(ctx context.Context, id int) (*StaffLanguageCode, error)
	Store(ctx context.Context, u *StaffLanguageCode) error
}
