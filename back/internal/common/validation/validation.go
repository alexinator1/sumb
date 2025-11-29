package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
)


// ValidationErrorHandler middleware обрабатывает ошибки валидации и возвращает стандартизированный ответ
// Должен быть зарегистрирован ПОСЛЕ всех хэндлеров, но ДО recovery middleware
func ValidationErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Проверяем, есть ли ошибки валидации только если ответ еще не был отправлен
		if len(c.Errors) > 0 && !c.Writer.Written() {
			var validationDetails []generated.ValidationErrorDetail
			var hasValidationError bool

			for _, err := range c.Errors {
				// Проверяем, является ли это ошибкой валидации от go-playground/validator
				if validationErr, ok := err.Err.(validator.ValidationErrors); ok {
					hasValidationError = true
					for _, fieldErr := range validationErr {
						validationDetails = append(validationDetails, generated.ValidationErrorDetail{
							Field:   fieldErr.Field(),
							Message: getValidationErrorMessage(fieldErr),
						})
					}
				} else if err.Type == gin.ErrorTypeBind {
					// Ошибка биндинга (например, невалидный JSON)
					hasValidationError = true
					validationDetails = append(validationDetails, generated.ValidationErrorDetail{
						Field:   "request",
						Message: err.Error(),
					})
				}
			}

			if hasValidationError {
				c.JSON(http.StatusBadRequest, generated.ValidationErrorResponse{
					Error:   "Validation Error",
					Message: "Request validation failed",
					Details: validationDetails,
				})
				c.Abort()
				return
			}
		}
	}
}

// getValidationErrorMessage возвращает понятное сообщение об ошибке валидации
func getValidationErrorMessage(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "required":
		return "Field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "len":
		return "Invalid length"
	case "oneof":
		return "Value must be one of the allowed values"
	default:
		return "Invalid value"
	}
}
