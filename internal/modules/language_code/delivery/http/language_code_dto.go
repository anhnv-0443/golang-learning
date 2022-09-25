package http

import (
	"go-app/internal/domain"
	"time"
)

// LanguageCodesResponse is array of LanguageCode response
type LanguageCodesResponse []LanguageCodeResponse

// LanguageCodeRequest is struct used for create st
type LanguageCodeRequest struct {
	Name        string
	Description string
}

// LanguageCodeResponse is struct used for LanguageCode
type LanguageCodeResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
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
	Data   *LanguageCodeRequest
}

// ConvertLanguageCodeToResponse DTO
func ConvertLanguageCodeToResponse(languageCodes *domain.LanguageCode) LanguageCodeResponse {
	return LanguageCodeResponse{
		ID:          int(languageCodes.ID),
		Name:        languageCodes.Name,
		Description: languageCodes.Description,
		CreatedAt:   languageCodes.CreatedAt,
		UpdatedAt:   languageCodes.UpdatedAt,
	}
}

//// ConvertLanguageCodesToResponse DTO
func ConvertLanguageCodesToResponse(LanguageCodes []domain.LanguageCode) LanguageCodesResponse {
	languageCodesRes := make(LanguageCodesResponse, 0)

	for _, u := range LanguageCodes {
		languageCodeRes := LanguageCodeResponse{
			ID:          int(u.ID),
			Name:        u.Name,
			Description: u.Description,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		}

		languageCodesRes = append(languageCodesRes, languageCodeRes)
	}

	return languageCodesRes
}

func ConvertRequestToEntity(u *LanguageCodeRequest) *domain.LanguageCode {
	return &domain.LanguageCode{
		Name:        u.Name,
		Description: u.Description,
	}
}
