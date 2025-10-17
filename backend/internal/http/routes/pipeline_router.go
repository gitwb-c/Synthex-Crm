package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func pipelineRouter(r *gin.Engine) {
	pipelines := r.Group("/pipelines")
	handler := handlers.NewPipelineHandler()

	{
		pipelines.GET("/read", handler.Read)
		pipelines.POST("/create", handler.Create)
		pipelines.PATCH("/update", handler.Update)
		pipelines.DELETE("/delete", handler.Delete)
	}
}
