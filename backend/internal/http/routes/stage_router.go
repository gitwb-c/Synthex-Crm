package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func stageRouter(r *gin.Engine) {
	stage := r.Group("/stage")
	handler := handlers.NewStageHandler()

	{
		stage.GET("/read", handler.Read)
		stage.POST("/create", handler.Create)
		stage.PATCH("/update", handler.Update)
		stage.DELETE("/delete", handler.Delete)
	}
}
