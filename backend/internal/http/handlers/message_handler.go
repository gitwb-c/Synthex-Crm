package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	client *ent.Client
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		client: db.Client,
	}
}

func (h *MessageHandler) Read(c *gin.Context) {}

func (h *MessageHandler) Create(c *gin.Context) {}

func (h *MessageHandler) Update(c *gin.Context) {}

func (h *MessageHandler) Delete(c *gin.Context) {}
