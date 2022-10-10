package domain

import (
	"context"
	"math/big"
)

// Project entity
type Project struct {
	ID   big.Int `json:"id"`
	Name string  `json:"name"`
}

// ProjectRepository represent the Project's usecases
type ProjectRepository interface {
	Fetch(context.Context) ([]Project, error)
	Find(ctx context.Context, id int) (*Project, error)
	Store(ctx context.Context, u *Project) error
}

// ProjectUseCase represent the Project's repository contract
type ProjectUseCase interface {
	Fetch(context.Context) ([]Project, error)
	Find(ctx context.Context, id int) (*Project, error)
	Store(ctx context.Context, u *Project) error
}
