package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	client *ent.Client
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{
		client: db.Client,
	}
}

func (h *EmployeeHandler) Read(c *gin.Context) {}

func (h *EmployeeHandler) Create(c *gin.Context) {}

func (h *EmployeeHandler) Update(c *gin.Context) {}

func (h *EmployeeHandler) Delete(c *gin.Context) {}
