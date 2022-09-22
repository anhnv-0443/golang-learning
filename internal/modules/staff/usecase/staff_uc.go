package usecase

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
)

// StaffUsecase ...
type StaffUsecase struct {
	repo domain.StaffRepository
}

// NewUsecase will create new StaffUsecase object
func NewUsecase(rp domain.StaffRepository) domain.StaffUsecase {
	return &StaffUsecase{
		repo: rp,
	}
}

// Fetch will fetch content from repo
func (uc *StaffUsecase) Fetch(c context.Context) ([]domain.Staff, error) {
	items, _ := uc.repo.Fetch(c)

	return items, nil
}

// Find will find content from repo
func (uc *StaffUsecase) Find(c context.Context, id int) (*domain.Staff, error) {
	item, err := uc.repo.Find(c, id)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return item, nil
}

// Store will create content from repo
func (uc *StaffUsecase) Store(c context.Context, staff *domain.Staff) error {
	if err := uc.repo.Store(c, staff); err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// Search is a function that returns list of Staffs filtered by query
func (uc *StaffUsecase) Search(ctx context.Context, q domain.StaffQueryParam) ([]domain.Staff, error) {
	items, err := uc.repo.Search(ctx, q)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return items, nil
}

// FindByQuery is a function that returns a staff filtered by query
func (uc *StaffUsecase) FindByQuery(ctx context.Context, q domain.StaffQueryParam) (*domain.Staff, error) {
	item, err := uc.repo.FindByQuery(ctx, q)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return item, nil
}
