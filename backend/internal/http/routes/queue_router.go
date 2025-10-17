package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func queueRouter(r *gin.Engine) {
	queue := r.Group("/queue")
	handler := handlers.NewQueueHandler()

	{
		queue.GET("/read", handler.Read)
		queue.POST("/create", handler.Create)
		queue.PATCH("/update", handler.Update)
		queue.DELETE("/delete", handler.Delete)
	}
}
