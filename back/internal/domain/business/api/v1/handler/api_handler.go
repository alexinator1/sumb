package handler

import (
	"net/http"
	"strconv"

	"github.com/alexinator1/sumb/back/internal/common/validation"
	adapter "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/adapter"
	"github.com/alexinator1/sumb/back/internal/domain/business/api/v1/builder"
	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
	businessvalidator "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/validator"
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
		validation.HandleValidationError(c, err)
		return 
	}

	// Валидация запроса
	validationDetails := businessvalidator.ValidateCreateBusinessRequest(req)
	if len(validationDetails) > 0 {
		c.JSON(http.StatusBadRequest, generated.ValidationErrorResponse{
			Error:   "Validation Error",
			Message: "Request validation failed",
			Details: validationDetails,
		})
		return
	}

	businessEntity, err := adapter.RequestToDomainBusiness(req)
	if err != nil {
		if validation.HandleValidationError(c, err) {
			return
		}
	}

	ownerEntity, err := adapter.RequestToDomainEmployee(req)
	if err != nil {
		if validation.HandleValidationError(c, err) {
			return
		}
	}

	createdBusiness, err := h.service.CreateBusinessWithOwner(c.Request.Context(), businessEntity, ownerEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, builder.BuildCreateResponse(createdBusiness))
}

func (h *ApiHandler) UpdateBusiness(c *gin.Context) {
	// TODO: add when service supports update
	c.Status(http.StatusNotImplemented)
}
