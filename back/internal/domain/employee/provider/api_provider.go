package provider

import (
	handler "github.com/alexinator1/sumb/back/internal/domain/employee/api/v1/handler"
	repo "github.com/alexinator1/sumb/back/internal/domain/employee/repository"
	service "github.com/alexinator1/sumb/back/internal/domain/employee/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeApiProvider struct {
	db *gorm.DB
}

func NewEmployeeApiProvider(db *gorm.DB) *EmployeeApiProvider {
	return &EmployeeApiProvider{db: db}
}

func (p *EmployeeApiProvider) AddApiV1Routes(router *gin.Engine) *gin.RouterGroup {
	handler := p.buildHandler()

	employeesRoutes := router.Group("/employees")
	{
		employeesRoutes.GET("/:id", handler.GetEmployee)
	}

	return employeesRoutes
}

func (p *EmployeeApiProvider) buildHandler() *handler.Handler {
	repo := repo.NewRepo(p.db)
	service := service.NewService(repo)

	return handler.NewHandler(service)
}
