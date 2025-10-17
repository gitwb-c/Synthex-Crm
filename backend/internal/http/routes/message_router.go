package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func messageRouter(r *gin.Engine) {
	message := r.Group("/message")

	messageService := services.NewMessageService()
	handler := handlers.NewMessageHandler(messageService)
	{
		message.GET("/read", handler.Read)
		message.POST("/create", handler.Create)
		message.PATCH("/update", handler.Update)
		message.DELETE("/delete", handler.Delete)
	}
}
