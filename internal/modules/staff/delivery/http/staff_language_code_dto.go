package http

import (
	"go-app/internal/domain"
)

type StaffLanguageCodesResponse []StlaffLanguageCodesResponse

type StaffLanguageCodesRequest struct {
	StaffId        int `json:"staff_id"`
	LanguageCodeId int `json:"language_code_id"`
}

type StlaffLanguageCodesResponse struct {
	StaffId        int `json:"staff_id"`
	LanguageCodeId int `json:"language_code_id"`
}

func ConvertStaffLanguageCodeRequestToEntity(u *StaffLanguageCodesRequest) *domain.StaffLanguageCode {
	return &domain.StaffLanguageCode{
		LanguageCodeId: u.LanguageCodeId,
		StaffId:        u.StaffId,
	}
}
