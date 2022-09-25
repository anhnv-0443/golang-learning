package domain

import (
	"context"
	"time"
)

// LanguageCode entity
type LanguageCode struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// LanguageCodeRepository represent the LanguageCode's usecases
type LanguageCodeRepository interface {
	Fetch(context.Context) ([]LanguageCode, error)
	Find(ctx context.Context, id int) (*LanguageCode, error)
	Store(ctx context.Context, u *LanguageCode) error
}

// LanguageCodeUseCase represent the LanguageCode's repository contract
type LanguageCodeUsecase interface {
	Fetch(context.Context) ([]LanguageCode, error)
	Find(ctx context.Context, id int) (*LanguageCode, error)
	Store(ctx context.Context, u *LanguageCode) error
}
