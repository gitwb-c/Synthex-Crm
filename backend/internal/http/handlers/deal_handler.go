package handlers

import (
	"net/http"

	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type DealHandler struct {
	service *services.DealService
}

func NewDealHandler(service *services.DealService) *DealHandler {
	return &DealHandler{service: service}
}

func (h *DealHandler) Read(c *gin.Context) {
	var req dto.ReadDealReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *DealHandler) Create(c *gin.Context) {
	var req dto.CreateDealReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *DealHandler) Update(c *gin.Context) {
	var req dto.UpdateDealReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *DealHandler) Delete(c *gin.Context) {
	var req dto.DeleteDealReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
