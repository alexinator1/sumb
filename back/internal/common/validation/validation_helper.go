package validation

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ValidateRequiredFields проверяет обязательные поля в структуре на основе OpenAPI required полей
func ValidateRequiredFields(req interface{}, requiredFields []string) error {
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct, got %s", v.Kind())
	}

	var missingFields []string
	for _, fieldName := range requiredFields {
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			continue
		}

		// Проверяем, является ли поле пустым (zero value)
		if isEmptyValue(field) {
			missingFields = append(missingFields, fieldName)
		}
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("required fields missing: %s", strings.Join(missingFields, ", "))
	}
	return nil
}

// ValidateEmailFormat проверяет формат email
func ValidateEmailFormat(email interface{}) error {
	var emailStr string

	// Обрабатываем разные типы email
	switch v := email.(type) {
	case string:
		emailStr = v
	case fmt.Stringer:
		emailStr = v.String()
	default:
		emailStr = fmt.Sprintf("%v", email)
	}

	if emailStr == "" {
		return nil // Пустой email уже проверен через ValidateRequiredFields
	}

	// Простая проверка формата email
	if !strings.Contains(emailStr, "@") || !strings.Contains(emailStr, ".") {
		return fmt.Errorf("invalid email format: %s", emailStr)
	}

	parts := strings.Split(emailStr, "@")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return fmt.Errorf("invalid email format: %s", emailStr)
	}

	// Проверяем, что после @ есть хотя бы одна точка
	if !strings.Contains(parts[1], ".") {
		return fmt.Errorf("invalid email format: %s", emailStr)
	}

	return nil
}

// isEmptyValue проверяет, является ли значение пустым (zero value)
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map:
		return v.IsNil()
	default:
		return false
	}
}

// HandleValidationError обрабатывает ошибку валидации и возвращает стандартизированный ответ
// Используйте эту функцию в хэндлерах для обработки ошибок валидации
func HandleValidationError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	validationDetails := make([]generated.ValidationErrorDetail, 0)

	// Проверяем, является ли это ошибкой валидации от go-playground/validator
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErr {
			validationDetails = append(validationDetails, generated.ValidationErrorDetail{
				Field:   fieldErr.Field(),
				Message: getValidationErrorMessage(fieldErr),
			})
		}
	} else {
		// Ошибка биндинга (например, невалидный JSON) или ошибка валидации обязательных полей
		// Парсим сообщение об ошибке для извлечения информации о полях
		errMsg := err.Error()
		if strings.Contains(errMsg, "required fields missing:") {
			// Извлекаем список отсутствующих полей
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
		} else if strings.Contains(errMsg, "email") && strings.Contains(errMsg, "failed to pass regex validation") {
			// Ошибка валидации формата email от ShouldBindJSON (openapi_types.Email)
			// Определяем поле по контексту запроса или используем дефолтное имя
			// В большинстве случаев это будет OwnerEmail для CreateBusinessRequest
			fieldName := extractEmailFieldName(c, errMsg)
			validationDetails = append(validationDetails, generated.ValidationErrorDetail{
				Field:   fieldName,
				Message: "Invalid email format",
			})
		} else if strings.Contains(errMsg, "invalid email format:") {
			// Ошибка валидации формата email от нашего валидатора
			validationDetails = append(validationDetails, generated.ValidationErrorDetail{
				Field:   "OwnerEmail",
				Message: "Invalid email format",
			})
		} else {
			// Общая ошибка биндинга
			validationDetails = append(validationDetails, generated.ValidationErrorDetail{
				Field:   "request",
				Message: errMsg,
			})
		}
	}

	c.JSON(http.StatusBadRequest, generated.ValidationErrorResponse{
		Error:   "Validation Error",
		Message: "Request validation failed",
		Details: validationDetails,
	})
	return true
}

// extractEmailFieldName пытается определить имя поля email из контекста ошибки
// По умолчанию возвращает "OwnerEmail" для CreateBusinessRequest
func extractEmailFieldName(c *gin.Context, errMsg string) string {
	// Пытаемся извлечь имя поля из сообщения об ошибке
	if strings.Contains(errMsg, "OwnerEmail") {
		return "OwnerEmail"
	} else if strings.Contains(errMsg, "ownerEmail") {
		return "OwnerEmail"
	} else if strings.Contains(errMsg, "owner_email") {
		return "OwnerEmail"
	}

	// По умолчанию для CreateBusinessRequest это OwnerEmail
	// Можно улучшить, анализируя путь запроса или другие признаки
	return "OwnerEmail"
}
