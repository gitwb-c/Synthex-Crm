package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func dealRouter(r *gin.Engine) {
	deal := r.Group("/deal")

	dealService := services.NewDealService()
	handler := handlers.NewDealHandler(dealService)
	{
		deal.GET("/read", handler.Read)
		deal.POST("/create", handler.Create)
		deal.PATCH("/update", handler.Update)
		deal.DELETE("/delete", handler.Delete)
	}
}
