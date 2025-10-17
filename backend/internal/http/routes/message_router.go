package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func messageRouter(r *gin.Engine) {
	message := r.Group("/message")
	handler := handlers.NewMessageHandler()

	{
		message.GET("/read", handler.Read)
		message.POST("/create", handler.Create)
		message.PATCH("/update", handler.Update)
		message.DELETE("/delete", handler.Delete)
	}
}
