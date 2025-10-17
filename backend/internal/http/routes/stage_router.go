package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func stageRouter(r *gin.Engine) {
	stage := r.Group("/stage")

	stageService := services.NewStageService()
	handler := handlers.NewStageHandler(stageService)
	{
		stage.GET("/read", handler.Read)
		stage.POST("/create", handler.Create)
		stage.PATCH("/update", handler.Update)
		stage.DELETE("/delete", handler.Delete)
	}
}
