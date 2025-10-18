package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func formRouter(r *gin.Engine) {
	form := r.Group("/form")

	handler := internal.InitializeFormHandler()
	{
		form.GET("/", handler.Read)
		form.POST("/", handler.Create)
		form.PATCH("/", handler.Update)
		form.DELETE("/", handler.Delete)
	}
}
