package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func queueRouter(r *gin.Engine) {
	queue := r.Group("/queue")

	queueService := services.NewQueueService()
	handler := handlers.NewQueueHandler(queueService)
	{
		queue.GET("/read", handler.Read)
		queue.POST("/create", handler.Create)
		queue.PATCH("/update", handler.Update)
		queue.DELETE("/delete", handler.Delete)
	}
}
