// internal/domain/employee/api/v1/handler.go
package handler

import (
	"net/http"
	"strconv"

	adapter "github.com/alexinator1/sumb/back/internal/domain/employee/api/v1/adapter"
	generated "github.com/alexinator1/sumb/back/internal/domain/employee/api/v1/employeegenerated"
	service "github.com/alexinator1/sumb/back/internal/domain/employee/service"
	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	service *service.EmployeeService
}

func NewHandler(service *service.EmployeeService) *ApiHandler {
	return &ApiHandler{service: service}
}

func (h *ApiHandler) GetEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	employee, err := h.service.GetEmployeeByID(c.Request.Context(), uint64(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, adapter.DomainToGeneratedEmployee(employee))
}

func (h *ApiHandler) CreateEmployee(c *gin.Context) {
	var req generated.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ent, err := adapter.GeneratedToDomainEmployee(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: when service has Create/Save, call it here and return persisted entity
	c.JSON(http.StatusOK, adapter.DomainToGeneratedEmployee(ent))
}

func (h *ApiHandler) DeleteEmployee(c *gin.Context) {
	// вызов service.Delete()
}
