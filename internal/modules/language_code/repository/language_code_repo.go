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
