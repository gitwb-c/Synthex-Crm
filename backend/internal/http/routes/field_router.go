package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func fieldRouter(r *gin.Engine) {
	field := r.Group("/field")
	handler := handlers.NewFieldHandler()

	{
		field.GET("/read", handler.Read)
		field.POST("/create", handler.Create)
		field.PATCH("/update", handler.Update)
		field.DELETE("/delete", handler.Delete)
	}
}
