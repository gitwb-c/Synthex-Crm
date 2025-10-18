package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func pipelineRouter(r *gin.Engine) {
	pipelines := r.Group("/pipelines")

	handler := internal.InitializePipelineHandler()
	{
		pipelines.GET("/", handler.Read)
		pipelines.POST("/", handler.Create)
		pipelines.PATCH("/", handler.Update)
		pipelines.DELETE("/", handler.Delete)
	}
}
