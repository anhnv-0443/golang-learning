package repository

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"gorm.io/gorm"
)

// StaffRepository ...
type StaffRepository struct {
	*gorm.DB
}

// NewRepository will create new postgres object
func NewRepository(db *gorm.DB) domain.StaffRepository {
	return &StaffRepository{DB: db}
}

// Fetch will fetch content from db
func (rp *StaffRepository) Fetch(c context.Context) ([]domain.Staff, error) {
	staffs := []domain.Staff{}
	if err := rp.DB.Find(&staffs).Error; err != nil {
		return staffs, errors.Wrap(err)
	}

	return staffs, nil
}

// Find will find content from db
func (rp *StaffRepository) Find(c context.Context, id int) (*domain.Staff, error) {
	staff := domain.Staff{}
	if err := rp.DB.Where("id = ?", id).First(&staff).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return &staff, nil
}

// Store will create data to db
func (rp *StaffRepository) Store(c context.Context, staff *domain.Staff) error {
	//passHash, err := utils.GeneratePassword(staff.Password)
	//if err != nil {
	//	return errors.Wrap(err)
	//}
	//staff.Password = passHash
	if err := rp.DB.Create(staff).Error; err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// Search is a function that returns a list of Staffs filtered by query
func (rp *StaffRepository) Search(ctx context.Context, q domain.StaffQueryParam) ([]domain.Staff, error) {
	var staffs []domain.Staff
	if err := rp.DB.Where("email = ? ", q.Email).Find(&staffs).Error; err != nil {
		return staffs, errors.Wrap(err)
	}

	return staffs, nil
}

// FindByQuery is a function that returns a Staffs filtered by query
func (rp *StaffRepository) FindByQuery(ctx context.Context, q domain.StaffQueryParam) (*domain.Staff, error) {
	staff := &domain.Staff{}
	if err := rp.DB.Where("email = ? ", q.Email).First(&staff).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return staff, nil
}
