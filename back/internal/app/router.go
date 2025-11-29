package app

import (
	"github.com/alexinator1/sumb/back/internal/common/validation"
	"github.com/gin-gonic/gin"
)

func buildRouter(appProviders *AppProviders) *gin.Engine {
	router := gin.Default()

	// Добавляем middleware для обработки ошибок валидации
	router.Use(validation.ValidationErrorHandler())

	// add employee routes
	appProviders.EmployeeProvider().AddApiV1Routes(router)

	// add business routes
	appProviders.BusinessProvider().AddApiV1Routes(router)

	return router
}
