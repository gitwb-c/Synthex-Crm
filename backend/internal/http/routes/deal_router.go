package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func dealRouter(r *gin.Engine) {
	deal := r.Group("/deal")
	handler := handlers.NewDealHandler()

	{
		deal.GET("/read", handler.Read)
		deal.POST("/create", handler.Create)
		deal.PATCH("/update", handler.Update)
		deal.DELETE("/delete", handler.Delete)
	}
}
