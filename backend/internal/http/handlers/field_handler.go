package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type FieldHandler struct {
	client *ent.Client
}

func NewFieldHandler() *FieldHandler {
	return &FieldHandler{
		client: db.Client,
	}
}

func (h *FieldHandler) Read(c *gin.Context) {}

func (h *FieldHandler) Create(c *gin.Context) {}

func (h *FieldHandler) Update(c *gin.Context) {}

func (h *FieldHandler) Delete(c *gin.Context) {}
