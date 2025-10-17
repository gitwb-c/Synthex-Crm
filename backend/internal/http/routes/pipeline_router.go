package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func pipelineRouter(r *gin.Engine) {
	pipelines := r.Group("/pipelines")

	pipelineService := services.NewPipelineService()
	handler := handlers.NewPipelineHandler(pipelineService)
	{
		pipelines.GET("/read", handler.Read)
		pipelines.POST("/create", handler.Create)
		pipelines.PATCH("/update", handler.Update)
		pipelines.DELETE("/delete", handler.Delete)
	}
}
