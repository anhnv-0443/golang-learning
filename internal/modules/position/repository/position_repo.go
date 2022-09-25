package repository

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"gorm.io/gorm"
)

// PositionRepository ...
type PositionRepository struct {
	*gorm.DB
}

// NewRepository will create new postgres object
func NewRepository(db *gorm.DB) domain.PositionRepository {
	return &PositionRepository{DB: db}
}

// Store will create data to db
func (rp *PositionRepository) Store(c context.Context, role *domain.Position) error {
	if err := rp.DB.Create(role).Error; err != nil {
		return errors.Wrap(err)
	}

	return nil
}
