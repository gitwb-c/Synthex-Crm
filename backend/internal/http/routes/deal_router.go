package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func dealRouter(r *gin.Engine) {
	deal := r.Group("/deal")

	handler := internal.InitializeDealHandler()
	{
		deal.GET("/", handler.Read)
		deal.POST("/", handler.Create)
		deal.PATCH("/", handler.Update)
		deal.DELETE("/", handler.Delete)
	}
}
