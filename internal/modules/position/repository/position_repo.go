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

func (rp *PositionRepository) Fetch(ctx context.Context) ([]domain.Position, error) {
	positions := []domain.Position{}
	if err := rp.DB.Find(&positions).Error; err != nil {
		return positions, errors.Wrap(err)
	}

	return positions, nil
}

func (rp *PositionRepository) Find(ctx context.Context, id int) (*domain.Position, error) {
	positions := domain.Position{}
	if err := rp.DB.Where("id = ?", id).First(&positions).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return &positions, nil
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
