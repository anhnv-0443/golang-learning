package repository

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"gorm.io/gorm"
)

// LanguageCodeRepository ...
type LanguageCodeRepository struct {
	*gorm.DB
}

func (rp *LanguageCodeRepository) Fetch(ctx context.Context) ([]domain.LanguageCode, error) {
	var languages []domain.LanguageCode
	if err := rp.DB.Find(&languages).Error; err != nil {
		return languages, errors.Wrap(err)
	}

	return languages, nil
}

func (rp *LanguageCodeRepository) Find(ctx context.Context, id int) (*domain.LanguageCode, error) {
	languages := domain.LanguageCode{}
	if err := rp.DB.Where("id = ?", id).First(&languages).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return &languages, nil
}

// NewRepository will create new postgres object
func NewRepository(db *gorm.DB) domain.LanguageCodeRepository {
	return &LanguageCodeRepository{DB: db}
}

// Store will create data to db
func (rp *LanguageCodeRepository) Store(c context.Context, role *domain.LanguageCode) error {
	if err := rp.DB.Create(role).Error; err != nil {
		return errors.Wrap(err)
	}

	return nil
}
