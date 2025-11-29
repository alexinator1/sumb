package provider

import (
	handler "github.com/alexinator1/sumb/back/internal/domain/employee/api/v1/handler"
	repo "github.com/alexinator1/sumb/back/internal/domain/employee/repository"
	service "github.com/alexinator1/sumb/back/internal/domain/employee/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeProvider struct {
	db *gorm.DB
	repo *repo.EmployeeRepo
}

func NewEmployeeProvider(db *gorm.DB) *EmployeeProvider {
	repo := repo.NewRepo(db)
	return &EmployeeProvider{db: db, repo: repo}
}

func (p *EmployeeProvider) AddApiV1Routes(router *gin.Engine) *gin.RouterGroup {
	h := p.buildHandler()

	employeesRoutes := router.Group("/employees")
	{
		employeesRoutes.GET("/:id", h.GetEmployee)
		employeesRoutes.POST("", h.CreateEmployee)
	}

	return employeesRoutes
}

func (p *EmployeeProvider) EmployeeRepo() *repo.EmployeeRepo {
	return p.repo
}

func (p *EmployeeProvider) buildHandler() *handler.ApiHandler {
	service := service.NewService(p.repo)

	return handler.NewHandler(service)
}
