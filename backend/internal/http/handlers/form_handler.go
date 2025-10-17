package handlers

import (
	"net/http"

	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type FormHandler struct {
	service *services.FormService
}

func NewFormHandler(service *services.FormService) *FormHandler {
	return &FormHandler{service: service}
}

func (h *FormHandler) Read(c *gin.Context) {
	var req dto.ReadFormReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *FormHandler) Create(c *gin.Context) {
	var req dto.CreateFormReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *FormHandler) Update(c *gin.Context) {
	var req dto.UpdateFormReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *FormHandler) Delete(c *gin.Context) {
	var req dto.DeleteFormReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
