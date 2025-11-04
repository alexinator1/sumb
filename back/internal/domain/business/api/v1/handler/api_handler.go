package handler

import (
	"net/http"
	"strconv"

	adapter "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/adapter"
	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
	service "github.com/alexinator1/sumb/back/internal/domain/business/service"
	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	service *service.BusinessService
}

func NewHandler(service *service.BusinessService) *ApiHandler {
	return &ApiHandler{service: service}
}

func (h *ApiHandler) GetBusiness(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	b, err := h.service.GetBusinessByID(c.Request.Context(), uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adapter.DomainToGeneratedBusiness(b))
}

func (h *ApiHandler) CreateBusiness(c *gin.Context) {
	var req generated.CreateBusinessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := generated.AssertCreateBusinessRequestRequired(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := generated.AssertCreateBusinessRequestConstraints(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	businessEntity, err := adapter.RequestToDomainBusiness(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ownerEntity, err := adapter.RequestToDomainEmployee(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.service.CreateBusinessWithOwner(c.Request.Context(), businessEntity, ownerEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, adapter.DomainToGeneratedBusiness(created))
}

func (h *ApiHandler) UpdateBusiness(c *gin.Context) {
	// TODO: add when service supports update
	c.Status(http.StatusNotImplemented)
}
