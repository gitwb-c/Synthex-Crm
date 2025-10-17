package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func fieldRouter(r *gin.Engine) {
	field := r.Group("/field")

	fieldService := services.NewFieldService()
	handler := handlers.NewFieldHandler(fieldService)
	{
		field.GET("/read", handler.Read)
		field.POST("/create", handler.Create)
		field.PATCH("/update", handler.Update)
		field.DELETE("/delete", handler.Delete)
	}
}
