package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type PipelineHandler struct {
	client *ent.Client
}

func NewPipelineHandler() *PipelineHandler {
	return &PipelineHandler{
		client: db.Client,
	}
}

func (h *PipelineHandler) Read(c *gin.Context) {}

func (h *PipelineHandler) Create(c *gin.Context) {}

func (h *PipelineHandler) Update(c *gin.Context) {}

func (h *PipelineHandler) Delete(c *gin.Context) {}
