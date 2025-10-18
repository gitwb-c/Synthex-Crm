package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func costumerRouter(r *gin.Engine) {
	costumer := r.Group("/costumer")

	handler := internal.InitializeCostumerHandler()
	{
		costumer.GET("/", handler.Read)
		costumer.POST("/", handler.Create)
		costumer.PATCH("/", handler.Update)
		costumer.DELETE("/", handler.Delete)
	}
}
