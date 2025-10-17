package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type QueueParticipantHandler struct {
	client *ent.Client
}

func NewQueueParticipantHandler() *QueueParticipantHandler {
	return &QueueParticipantHandler{
		client: db.Client,
	}
}

func (h *QueueParticipantHandler) Read(c *gin.Context) {}

func (h *QueueParticipantHandler) Create(c *gin.Context) {}

func (h *QueueParticipantHandler) Update(c *gin.Context) {}

func (h *QueueParticipantHandler) Delete(c *gin.Context) {}
