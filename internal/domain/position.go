package domain

import (
	"context"
	"time"
)

// Position entity
type Position struct {
	ID				uint       `json:"id"`
	Name			string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type PositionRepository interface {
	Store(ctx context.Context, u *Position) error
}