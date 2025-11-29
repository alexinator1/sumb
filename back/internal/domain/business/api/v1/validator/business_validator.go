package validator

import (
	"strings"

	"github.com/alexinator1/sumb/back/internal/common/validation"
	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
)

// ValidateCreateBusinessRequest валидирует запрос на создание бизнеса
// Возвращает ValidationError с деталями ошибок или nil, если валидация прошла успешно
func ValidateCreateBusinessRequest(req generated.CreateBusinessRequest) []generated.ValidationErrorDetail {
	var validationDetails []generated.ValidationErrorDetail

	// Валидация обязательных полей согласно OpenAPI спецификации
	requiredFields := []string{"Name", "OwnerFirstName", "OwnerLastName", "OwnerEmail", "OwnerPhone"}
	if err := validation.ValidateRequiredFields(req, requiredFields); err != nil {
		// Парсим ошибку для извлечения списка отсутствующих полей
		errMsg := err.Error()
		if strings.Contains(errMsg, "required fields missing:") {
			parts := strings.Split(errMsg, "required fields missing: ")
			if len(parts) > 1 {
				fields := strings.Split(strings.TrimSpace(parts[1]), ", ")
				for _, field := range fields {
					validationDetails = append(validationDetails, generated.ValidationErrorDetail{
						Field:   field,
						Message: "Field is required",
					})
				}
			}
		}
	}

	// Валидация формата email
		if err := validation.ValidateEmailFormat(req.OwnerEmail); err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "invalid email format:") {
			validationDetails = append(validationDetails, generated.ValidationErrorDetail{
				Field:   "OwnerEmail",
				Message: "Invalid email format",
			})
		}
	}

	// Валидация пароля
	if req.Password == "" {
		validationDetails = append(validationDetails, generated.ValidationErrorDetail{
			Field:   "Password",
			Message: "Password is required",
		})
	}
	if req.PasswordConfirmation == "" || req.PasswordConfirmation != req.Password {
			validationDetails = append(validationDetails, generated.ValidationErrorDetail{
			Field:   "PasswordConfirmation",
			Message: "Password confirmation does not match",
		})
	}
	if len(validationDetails) > 0 {
		return validationDetails
	}
	return []generated.ValidationErrorDetail{}
}
