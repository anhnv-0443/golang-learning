package http

import (
	"go-app/internal/domain"
	"time"
)

// StaffsResponse is array of staff response
type StaffsResponse []StaffResponse

// StaffRequest is struct used for create staff
type StaffRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	StaffCode   string `json:"staff_code" validate:"required"`
	BirthDate   string `json:"birth_date" validate:"required"`
	Gender      int    `json:"gender" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	UnitID      int    `json:"unit_id" validate:"required"`
	StaffTypeID int    `json:"staff_type_id" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

// StaffResponse is struct used for staff
type StaffResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	StaffCode   string    `json:"staff_code"`
	BirthDate   string    `json:"birth_date"`
	Gender      int       `json:"gender"`
	PhoneNumber string    `json:"phone_number"`
	UnitID      int       `json:"unit_id"`
	StaffTypeID int       `json:"staff_type_id"`
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
		UnitID:      staffs.UnitID,
		StaffTypeID: staffs.StaffTypeID,
		Address:     staffs.Address,
		CreatedAt:   staffs.CreatedAt,
		UpdatedAt:   staffs.UpdatedAt,
	}
}

// ConvertStaffsToResponse DTO
func ConvertStaffsToResponse(staffs []domain.Staff) StaffsResponse {
	staffsRes := make(StaffsResponse, 0)

	for _, u := range staffs {
		staffRes := StaffResponse{
			ID:          u.ID,
			Name:        u.Name,
			Email:       u.Email,
			StaffCode:   u.StaffCode,
			BirthDate:   u.BirthDate,
			Gender:      u.Gender,
			PhoneNumber: u.PhoneNumber,
			UnitID:      u.UnitID,
			StaffTypeID: u.StaffTypeID,
			Address:     u.Address,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		}

		staffsRes = append(staffsRes, staffRes)
	}

	return staffsRes
}

// ConvertRequestToEntity DTO
func ConvertRequestToEntity(u *StaffRequest) *domain.Staff {
	return &domain.Staff{
		Name:        u.Name,
		Email:       u.Email,
		StaffCode:   u.StaffCode,
		BirthDate:   u.BirthDate,
		Gender:      u.Gender,
		PhoneNumber: u.PhoneNumber,
		UnitID:      u.UnitID,
		StaffTypeID: u.StaffTypeID,
		Address:     u.Address,
	}
}
