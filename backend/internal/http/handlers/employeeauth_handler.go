package handlers

import (
	"net/http"

	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type EmployeeAuthHandler struct {
	service *services.EmployeeAuthService
}

func NewEmployeeAuthHandler(service *services.EmployeeAuthService) *EmployeeAuthHandler {
	return &EmployeeAuthHandler{service: service}
}

func (h *EmployeeAuthHandler) Read(c *gin.Context) {
	var req dto.ReadEmployeeAuthReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *EmployeeAuthHandler) Create(c *gin.Context) {
	var req dto.CreateEmployeeAuthReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *EmployeeAuthHandler) Update(c *gin.Context) {
	var req dto.UpdateEmployeeAuthReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *EmployeeAuthHandler) Delete(c *gin.Context) {
	var req dto.DeleteEmployeeAuthReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
