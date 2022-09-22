package domain

import (
	"context"
)

// LanguageCode entity
type LanguageCode struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// LanguageCodeRepository represent the LanguageCode's usecases
type LanguageCodeRepository interface {
	Fetch(context.Context) ([]LanguageCode, error)
	Find(ctx context.Context, id int) (*LanguageCode, error)
	Store(ctx context.Context, u *LanguageCode) error
}

// LanguageCodeUseCase represent the LanguageCode's repository contract
type LanguageCodeUseCase interface {
	Fetch(context.Context) ([]LanguageCode, error)
	Find(ctx context.Context, id int) (*LanguageCode, error)
	Store(ctx context.Context, u *LanguageCode) error
}
