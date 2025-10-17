package handlers

import (
	"net/http"

	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type QueueHandler struct {
	service *services.QueueService
}

func NewQueueHandler(service *services.QueueService) *QueueHandler {
	return &QueueHandler{service: service}
}

func (h *QueueHandler) Read(c *gin.Context) {
	var req dto.ReadQueueReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *QueueHandler) Create(c *gin.Context) {
	var req dto.CreateQueueReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *QueueHandler) Update(c *gin.Context) {
	var req dto.UpdateQueueReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *QueueHandler) Delete(c *gin.Context) {
	var req dto.DeleteQueueReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
