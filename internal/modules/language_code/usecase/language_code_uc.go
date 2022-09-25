package usecase

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
)

// LanguageCodeUsecase ...
type LanguageCodeUsecase struct {
	repo domain.LanguageCodeRepository
}

// NewUsecase will create new LanguageCodeUsecase object
func NewUsecase(rp domain.LanguageCodeRepository) domain.LanguageCodeUsecase {
	return &LanguageCodeUsecase{
		repo: rp,
	}
}

// Fetch will fetch content from repo
func (uc *LanguageCodeUsecase) Fetch(c context.Context) ([]domain.LanguageCode, error) {
	items, _ := uc.repo.Fetch(c)

	return items, nil
}

// Find will find content from repo
func (uc *LanguageCodeUsecase) Find(c context.Context, id int) (*domain.LanguageCode, error) {
	item, err := uc.repo.Find(c, id)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return item, nil
}

// Store will create content from repo
func (uc *LanguageCodeUsecase) Store(c context.Context, languageCode *domain.LanguageCode) error {
	if err := uc.repo.Store(c, languageCode); err != nil {
		return errors.Wrap(err)
	}

	return nil
}
