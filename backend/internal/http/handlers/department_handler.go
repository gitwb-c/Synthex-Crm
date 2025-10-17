package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	client *ent.Client
}

func NewDepartmentHandler() *DepartmentHandler {
	return &DepartmentHandler{
		client: db.Client,
	}
}

func (h *DepartmentHandler) Read(c *gin.Context) {}

func (h *DepartmentHandler) Create(c *gin.Context) {}

func (h *DepartmentHandler) Update(c *gin.Context) {}

func (h *DepartmentHandler) Delete(c *gin.Context) {}
