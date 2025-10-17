package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type StageHandler struct {
	client *ent.Client
}

func NewStageHandler() *StageHandler {
	return &StageHandler{
		client: db.Client,
	}
}

func (h *StageHandler) Read(c *gin.Context) {}

func (h *StageHandler) Create(c *gin.Context) {}

func (h *StageHandler) Update(c *gin.Context) {}

func (h *StageHandler) Delete(c *gin.Context) {}
