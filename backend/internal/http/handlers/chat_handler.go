package handlers

import (
	"net/http"

	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	service *services.ChatService
}

func NewChatHandler(service *services.ChatService) *ChatHandler {
	return &ChatHandler{service: service}
}

// chat, err := h.service.Read(req)
// if err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// }

// c.JSON(http.StatusOK, chat)

func (h *ChatHandler) Read(c *gin.Context) {
	var req dto.ReadChatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func (h *ChatHandler) Create(c *gin.Context) {
	var req dto.CreateChatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *ChatHandler) Update(c *gin.Context) {
	var req dto.UpdateChatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *ChatHandler) Delete(c *gin.Context) {
	var req dto.DeleteChatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
