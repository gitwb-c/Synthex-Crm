package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func fieldRouter(r *gin.Engine) {
	field := r.Group("/field")

	handler := internal.InitializeFieldHandler()
	{
		field.GET("/", handler.Read)
		field.POST("/", handler.Create)
		field.PATCH("/", handler.Update)
		field.DELETE("/", handler.Delete)
	}
}
