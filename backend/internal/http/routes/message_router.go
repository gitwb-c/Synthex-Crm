package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func messageRouter(r *gin.Engine) {
	message := r.Group("/message")

	handler := internal.InitializeMessageHandler()
	{
		message.GET("/", handler.Read)
		message.POST("/", handler.Create)
		message.PATCH("/", handler.Update)
		message.DELETE("/", handler.Delete)
	}
}
