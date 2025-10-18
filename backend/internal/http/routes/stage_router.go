package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func stageRouter(r *gin.Engine) {
	stage := r.Group("/stage")

	handler := internal.InitializeStageHandler()
	{
		stage.GET("/", handler.Read)
		stage.POST("/", handler.Create)
		stage.PATCH("/", handler.Update)
		stage.DELETE("/", handler.Delete)
	}
}
