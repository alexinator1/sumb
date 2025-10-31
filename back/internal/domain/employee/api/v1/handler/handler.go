// internal/domain/employee/api/v1/handler.go
package handler

import (
	"net/http"

	"github.com/alexinator1/sumb/back/internal/domain/employee/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *employee.EmployeeService
}

func NewHandler(service *employee.EmployeeService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetEmployee(c *gin.Context) {
	employee, err := h.service.GetEmployeeByID(c.Request.Context(), 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	// вызов service.Save()
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	// вызов service.Delete()
}
