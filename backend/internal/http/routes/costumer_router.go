package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func costumerRouter(r *gin.Engine) {
	costumer := r.Group("/costumer")

	costumerService := services.NewCostumerService()
	handler := handlers.NewCostumerHandler(costumerService)
	{
		costumer.GET("/read", handler.Read)
		costumer.POST("/create", handler.Create)
		costumer.PATCH("/update", handler.Update)
		costumer.DELETE("/delete", handler.Delete)
	}
}
