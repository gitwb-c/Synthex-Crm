package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func formRouter(r *gin.Engine) {
	form := r.Group("/form")
	handler := handlers.NewFormHandler()

	{
		form.GET("/read", handler.Read)
		form.POST("/create", handler.Create)
		form.PATCH("/update", handler.Update)
		form.DELETE("/delete", handler.Delete)
	}
}
