package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func chatRouter(r *gin.Engine) {
	chat := r.Group("/chat")
	handler := handlers.NewChatHandler()

	{
		chat.GET("/read", handler.Read)
		chat.POST("/create", handler.Create)
		chat.PATCH("/update", handler.Update)
		chat.DELETE("/delete", handler.Delete)
	}
}
