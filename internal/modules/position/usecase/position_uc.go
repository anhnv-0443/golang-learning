package usecase

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
)

// PositionUsecase ...
type PositionUsecase struct {
	repo domain.PositionRepository
}

// NewUsecase will create new PositionUsecase object
func NewUsecase(rp domain.PositionRepository) domain.PositionUsecase {
	return &PositionUsecase{
		repo: rp,
	}
}

// Fetch will fetch content from repo
func (uc *PositionUsecase) Fetch(c context.Context) ([]domain.Position, error) {
	items, _ := uc.repo.Fetch(c)

	return items, nil
}

// Find will find content from repo
func (uc *PositionUsecase) Find(c context.Context, id int) (*domain.Position, error) {
	item, err := uc.repo.Find(c, id)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return item, nil
}

// Store will create content from repo
func (uc *PositionUsecase) Store(c context.Context, position *domain.Position) error {
	if err := uc.repo.Store(c, position); err != nil {
		return errors.Wrap(err)
	}

	return nil
}
