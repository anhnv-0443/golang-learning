package domain

import (
	"context"
	"time"
)

// LanguageCode entity
type LanguageCode struct {
	ID				uint       `json:"id"`
	Name			string     `json:"name"`
	Description		string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type LanguageCodeRepository interface {
	Store(ctx context.Context, u *LanguageCode) error
}