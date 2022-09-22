package http

import (
	"go-app/internal/domain"

	"math/big"
	"time"
)

// StaffsResponse is array of staff response
type StaffsResponse []StaffResponse

// StaffRequest is struct used for create st
type StaffRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	StaffCode   string `json:"staff_code" validate:"required"`
	BirthDate   string `json:"birth_date" validate:"required"`
	Gender      int    `json:"gender" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	UnitId      int    `json:"unit_id" validate:"required"`
	StaffTypeId int    `json:"staff_type_id" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

// StaffResponse is struct used for staff
type StaffResponse struct {
	ID          big.Int   `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	StaffCode   string    `json:"staff_code"`
	BirthDate   string    `json:"birth_date"`
	Gender      int       `json:"gender"`
	PhoneNumber string    `json:"phone_number"`
	UnitId      int       `json:"unit_id"`
	StaffTypeId int       `json:"staff_type_id"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ErrorResponse is struct when error
type ErrorResponse struct {
	Message string `json:"message"`
}

// StatusResponse is struct when success
type StatusResponse struct {
	Status bool `json:"status"`
}

// ConvertStaffToResponse DTO
func ConvertStaffToResponse(staffs *domain.Staff) StaffResponse {
	return StaffResponse{
		ID:          staffs.ID,
		Name:        staffs.Name,
		Email:       staffs.Email,
		StaffCode:   staffs.StaffCode,
		BirthDate:   staffs.BirthDate,
		Gender:      staffs.Gender,
		PhoneNumber: staffs.PhoneNumber,
		UnitId:      staffs.UnitId,
		StaffTypeId: staffs.StaffTypeId,
		Address:     staffs.Address,
		CreatedAt:   staffs.CreatedAt,
		UpdatedAt:   staffs.UpdatedAt,
	}
}

//// ConvertstaffsToResponse DTO
func ConvertStaffsToResponse(staffs []domain.Staff) StaffsResponse {
	staffsRes := make(StaffsResponse, 0)

	for _, u := range staffs {
		staffRes := StaffResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			StaffCode: u.StaffCode,
			BirthDate: u.BirthDate,
			//BirthDate:   u.BirthDate.Format("Y-m-d"),
			Gender:      u.Gender,
			PhoneNumber: u.PhoneNumber,
			UnitId:      u.UnitId,
			StaffTypeId: u.StaffTypeId,
			Address:     u.Address,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		}

		staffsRes = append(staffsRes, staffRes)
	}

	return staffsRes
}

//
//// ConvertRequestToEntity DTO
func ConvertRequestToEntity(u *StaffRequest) *domain.Staff {
	return &domain.Staff{
		Name:  u.Name,
		Email: u.Email,
	}
}
