package registry

import (
	"go-app/internal/domain"
	authUC "go-app/internal/modules/auth/usecase"
	roleUC "go-app/internal/modules/role/usecase"
	staffUC "go-app/internal/modules/staff/usecase"
	userUC "go-app/internal/modules/user/usecase"
)

// Usecase registry
type Usecase struct {
	RoleUsecase  domain.RoleUsecase
	UserUsecase  domain.UserUsecase
	AuthUsecase  domain.AuthUsecase
	StaffUsecase domain.StaffUsecase
}

// NewUsecase implements from interface
func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{
		RoleUsecase:  roleUC.NewUsecase(repo.RoleRepository),
		UserUsecase:  userUC.NewUsecase(repo.UserRepository),
		AuthUsecase:  authUC.NewUsecase(repo.UserRepository),
		StaffUsecase: staffUC.NewUsecase(repo.StaffRepository),
	}
}
