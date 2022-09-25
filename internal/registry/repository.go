package registry

import (
	"go-app/internal/domain"
	languageCodeRepo "go-app/internal/modules/language_code/repository"
	positionRepo "go-app/internal/modules/position/repository"
	roleRepo "go-app/internal/modules/role/repository"
	staffRepo "go-app/internal/modules/staff/repository"
	userRepo "go-app/internal/modules/user/repository"

	"gorm.io/gorm"
)

// Repository registry
type Repository struct {
	RoleRepository         domain.RoleRepository
	UserRepository         domain.UserRepository
	StaffRepository        domain.StaffRepository
	PositionRepository     domain.PositionRepository
	LanguageCodeRepository domain.LanguageCodeRepository
}

// NewRepository implements from interface
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		RoleRepository:         roleRepo.NewRepository(db),
		UserRepository:         userRepo.NewRepository(db),
		StaffRepository:        staffRepo.NewRepository(db),
		PositionRepository:     positionRepo.NewRepository(db),
		LanguageCodeRepository: languageCodeRepo.NewRepository(db),
	}
}
