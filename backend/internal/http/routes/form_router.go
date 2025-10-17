package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func formRouter(r *gin.Engine) {
	form := r.Group("/form")

	formService := services.NewFormService()
	handler := handlers.NewFormHandler(formService)
	{
		form.GET("/read", handler.Read)
		form.POST("/create", handler.Create)
		form.PATCH("/update", handler.Update)
		form.DELETE("/delete", handler.Delete)
	}
}
