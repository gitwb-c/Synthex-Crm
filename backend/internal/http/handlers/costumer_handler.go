package handlers

import (
	"net/http"

	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type CostumerHandler struct {
	service *services.CostumerService
}

func NewCostumerHandler(service *services.CostumerService) *CostumerHandler {
	return &CostumerHandler{service: service}
}

func (h *CostumerHandler) Read(c *gin.Context) {
	var req dto.ReadCostumerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *CostumerHandler) Create(c *gin.Context) {
	var req dto.CreateCostumerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *CostumerHandler) Update(c *gin.Context) {
	var req dto.UpdateCostumerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *CostumerHandler) Delete(c *gin.Context) {
	var req dto.DeleteCostumerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
