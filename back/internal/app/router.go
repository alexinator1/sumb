package app

import (
	"github.com/gin-gonic/gin"
)

func buildRouter(appProviders *AppProviders) *gin.Engine {
	router := gin.Default()
	_ = appProviders // keep parameter referenced until real handlers use it

	// add employee routes
	appProviders.EmployeeApiProvider().AddApiV1Routes(router)

	return router
}
