package provider

import (
	handler "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/handler"
	repo "github.com/alexinator1/sumb/back/internal/domain/business/repository"
	service "github.com/alexinator1/sumb/back/internal/domain/business/service"
	employeeRepo "github.com/alexinator1/sumb/back/internal/domain/employee/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BusinessProvider struct {
	db *gorm.DB
}

func NewBusinessProvider(db *gorm.DB) *BusinessProvider {
	return &BusinessProvider{db: db}
}

func (p *BusinessProvider) AddApiV1Routes(router *gin.Engine) *gin.RouterGroup {
	h := p.buildHandler()

	routes := router.Group("/business")
	{
		routes.GET("/:id", h.GetBusiness)
		routes.POST("", h.CreateBusiness)
		routes.PUT("/:id", h.UpdateBusiness)
	}
	return routes
}

func (p *BusinessProvider) buildHandler() *handler.ApiHandler {
	r := repo.NewRepo(p.db)
	empRepo := employeeRepo.NewRepo(p.db)
	s := service.NewService(r, empRepo, p.db)
	return handler.NewHandler(s)
}
