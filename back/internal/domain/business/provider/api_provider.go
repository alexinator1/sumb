package provider

import (
	handler "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/handler"
	repo "github.com/alexinator1/sumb/back/internal/domain/business/repository"
	service "github.com/alexinator1/sumb/back/internal/domain/business/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BusinessApiProvider struct {
	db *gorm.DB
}

func NewBusinessApiProvider(db *gorm.DB) *BusinessApiProvider {
	return &BusinessApiProvider{db: db}
}

func (p *BusinessApiProvider) AddApiV1Routes(router *gin.Engine) *gin.RouterGroup {
	h := p.buildHandler()

	routes := router.Group("/businesses")
	{
		routes.GET("/:id", h.GetBusiness)
		routes.POST("/", h.CreateBusiness)
		routes.PUT("/:id", h.UpdateBusiness)
	}
	return routes
}

func (p *BusinessApiProvider) buildHandler() *handler.ApiHandler {
	r := repo.NewRepo(p.db)
	s := service.NewService(r)
	return handler.NewHandler(s)
}
