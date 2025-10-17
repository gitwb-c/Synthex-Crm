package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	client *ent.Client
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		client: db.Client,
	}
}

func (h *ChatHandler) Read(c *gin.Context) {}

func (h *ChatHandler) Create(c *gin.Context) {}

func (h *ChatHandler) Update(c *gin.Context) {}

func (h *ChatHandler) Delete(c *gin.Context) {}
