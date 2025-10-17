package handlers

import (
	"net/http"

	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	service *services.DepartmentService
}

func NewDepartmentHandler(service *services.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: service}
}

func (h *DepartmentHandler) Read(c *gin.Context) {
	var req dto.ReadDepartmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *DepartmentHandler) Create(c *gin.Context) {
	var req dto.CreateDepartmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *DepartmentHandler) Update(c *gin.Context) {
	var req dto.UpdateDepartmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *DepartmentHandler) Delete(c *gin.Context) {
	var req dto.DeleteDepartmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
