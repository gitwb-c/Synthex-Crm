package handlers

import (
	"net/http"

	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type FieldHandler struct {
	service *services.FieldService
}

func NewFieldHandler(service *services.FieldService) *FieldHandler {
	return &FieldHandler{service: service}
}

func (h *FieldHandler) Read(c *gin.Context) {
	var req dto.ReadFieldReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *FieldHandler) Create(c *gin.Context) {
	var req dto.CreateFieldReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *FieldHandler) Update(c *gin.Context) {
	var req dto.UpdateFieldReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *FieldHandler) Delete(c *gin.Context) {
	var req dto.DeleteFieldReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
