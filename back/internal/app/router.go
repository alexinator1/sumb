package app

import "github.com/gin-gonic/gin"

func buildRouter(appProviders *AppProviders) *gin.Engine {
	router := gin.Default()
	_ = appProviders // keep parameter referenced until real handlers use it

	employees := router.Group("/employees")
	{
		// GET /employees/:employeeId
		employees.GET("/:employeeId", func(c *gin.Context) {
			employeeId := c.Param("employeeId")

			// Основa хэндлера для getEmployee.
			// TODO: заменить тело на реальную реализацию:
			// - валидация employeeId
			// - вызов слоя сервисов/контроллера из appProviders
			// - формирование ответа и обработка ошибок
			c.JSON(501, gin.H{
				"error":      "getEmployee not implemented",
				"employeeId": employeeId,
			})
		})
	}

	return router
}
