package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type DealHandler struct {
	client *ent.Client
}

func NewDealHandler() *DealHandler {
	return &DealHandler{
		client: db.Client,
	}
}

func (h *DealHandler) Read(c *gin.Context) {}

func (h *DealHandler) Create(c *gin.Context) {}

func (h *DealHandler) Update(c *gin.Context) {}

func (h *DealHandler) Delete(c *gin.Context) {}
