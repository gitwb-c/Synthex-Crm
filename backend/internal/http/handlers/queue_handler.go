package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type QueueHandler struct {
	client *ent.Client
}

func NewQueueHandler() *QueueHandler {
	return &QueueHandler{
		client: db.Client,
	}
}

func (h *QueueHandler) Read(c *gin.Context) {}

func (h *QueueHandler) Create(c *gin.Context) {}

func (h *QueueHandler) Update(c *gin.Context) {}

func (h *QueueHandler) Delete(c *gin.Context) {}
