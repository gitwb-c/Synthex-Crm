package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func chatRouter(r *gin.Engine) {
	chat := r.Group("/chat")

	chatService := services.NewChatService()
	handler := handlers.NewChatHandler(chatService)
	{
		chat.GET("/read", handler.Read)
		chat.POST("/create", handler.Create)
		chat.PATCH("/update", handler.Update)
		chat.DELETE("/delete", handler.Delete)
	}
}
