package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func costumerRouter(r *gin.Engine) {
	costumer := r.Group("/costumer")
	handler := handlers.NewCostumerHandler()

	{
		costumer.GET("/read", handler.Read)
		costumer.POST("/create", handler.Create)
		costumer.PATCH("/update", handler.Update)
		costumer.DELETE("/delete", handler.Delete)
	}
}
