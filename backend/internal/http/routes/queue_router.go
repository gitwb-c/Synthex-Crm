package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func queueRouter(r *gin.Engine) {
	queue := r.Group("/queue")

	handler := internal.InitializeQueueHandler()
	{
		queue.GET("/", handler.Read)
		queue.POST("/", handler.Create)
		queue.PATCH("/", handler.Update)
		queue.DELETE("/", handler.Delete)
	}
}
