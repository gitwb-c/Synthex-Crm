package handlers

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
	"github.com/gin-gonic/gin"
)

type EmployeeAuthHandler struct {
	client *ent.Client
}

func NewEmployeeAuthHandler() *EmployeeAuthHandler {
	return &EmployeeAuthHandler{
		client: db.Client,
	}
}

func (h *EmployeeAuthHandler) Read(c *gin.Context) {}

func (h *EmployeeAuthHandler) Create(c *gin.Context) {}

func (h *EmployeeAuthHandler) Update(c *gin.Context) {}

func (h *EmployeeAuthHandler) Delete(c *gin.Context) {}
