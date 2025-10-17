package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type FormHandler struct {
	client *ent.Client
}

func NewFormHandler() *FormHandler {
	return &FormHandler{
		client: db.Client,
	}
}

func (h *FormHandler) Read(c *gin.Context) {}

func (h *FormHandler) Create(c *gin.Context) {}

func (h *FormHandler) Update(c *gin.Context) {}

func (h *FormHandler) Delete(c *gin.Context) {}
