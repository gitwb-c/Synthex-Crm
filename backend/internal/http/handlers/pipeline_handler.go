package handlers

import (
	"net/http"

	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type PipelineHandler struct {
	service *services.PipelineService
}

func NewPipelineHandler(service *services.PipelineService) *PipelineHandler {
	return &PipelineHandler{service: service}
}

func (h *PipelineHandler) Read(c *gin.Context) {
	var req dto.ReadPipelineReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *PipelineHandler) Create(c *gin.Context) {
	var req dto.CreatePipelineReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *PipelineHandler) Update(c *gin.Context) {
	var req dto.UpdatePipelineReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *PipelineHandler) Delete(c *gin.Context) {
	var req dto.DeletePipelineReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
