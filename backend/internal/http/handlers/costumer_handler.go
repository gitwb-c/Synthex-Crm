package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type CostumerHandler struct {
	client *ent.Client
}

func NewCostumerHandler() *CostumerHandler {
	return &CostumerHandler{
		client: db.Client,
	}
}

func (h *CostumerHandler) Read(c *gin.Context) {}

func (h *CostumerHandler) Create(c *gin.Context) {}

func (h *CostumerHandler) Update(c *gin.Context) {}

func (h *CostumerHandler) Delete(c *gin.Context) {}
