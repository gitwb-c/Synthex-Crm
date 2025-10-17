package handlers

import (
	"net/http"

	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type StageHandler struct {
	service *services.StageService
}

func NewStageHandler(service *services.StageService) *StageHandler {
	return &StageHandler{service: service}
}

func (h *StageHandler) Read(c *gin.Context) {
	var req dto.ReadStageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *StageHandler) Create(c *gin.Context) {
	var req dto.CreateStageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *StageHandler) Update(c *gin.Context) {
	var req dto.UpdateStageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *StageHandler) Delete(c *gin.Context) {
	var req dto.DeleteStageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
