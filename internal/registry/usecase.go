package registry

import (
	"go-app/internal/domain"
	authUC "go-app/internal/modules/auth/usecase"
	languageCodeUC "go-app/internal/modules/language_code/usecase"
	positionUC "go-app/internal/modules/position/usecase"
	roleUC "go-app/internal/modules/role/usecase"
	staffUC "go-app/internal/modules/staff/usecase"
	userUC "go-app/internal/modules/user/usecase"
)

// Usecase registry
type Usecase struct {
	RoleUsecase         domain.RoleUsecase
	UserUsecase         domain.UserUsecase
	AuthUsecase         domain.AuthUsecase
	StaffUsecase        domain.StaffUsecase
	PositionUsecase     domain.PositionUsecase
	LanguageCodeUsecase domain.LanguageCodeUsecase
}

// NewUsecase implements from interface
func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{
		RoleUsecase:         roleUC.NewUsecase(repo.RoleRepository),
		UserUsecase:         userUC.NewUsecase(repo.UserRepository),
		AuthUsecase:         authUC.NewUsecase(repo.UserRepository),
		StaffUsecase:        staffUC.NewUsecase(repo.StaffRepository),
		PositionUsecase:     positionUC.NewUsecase(repo.PositionRepository),
		LanguageCodeUsecase: languageCodeUC.NewUsecase(repo.LanguageCodeRepository),
	}
}
