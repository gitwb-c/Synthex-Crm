package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func chatRouter(r *gin.Engine) {
	chat := r.Group("/chat")

	handler := internal.InitializeChatHandler()
	{
		chat.GET("/", handler.Read)
		chat.POST("/", handler.Create)
		chat.PATCH("/", handler.Update)
		chat.DELETE("/", handler.Delete)
	}
}
