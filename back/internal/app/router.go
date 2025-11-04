package app

import (
	"github.com/gin-gonic/gin"
)

func buildRouter(appProviders *AppProviders) *gin.Engine {
	router := gin.Default()

	// add employee routes
	appProviders.EmployeeApiProvider().AddApiV1Routes(router)

	// add business routes
	appProviders.BusinessApiProvider().AddApiV1Routes(router)

	return router
}
