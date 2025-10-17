package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func employeeAuthRouter(r *gin.Engine) {
	employeeAuth := r.Group("/employeeauth")
	handler := handlers.NewEmployeeAuthHandler()

	{
		employeeAuth.GET("/read", handler.Read)
		employeeAuth.POST("/create", handler.Create)
		employeeAuth.PATCH("/update", handler.Update)
		employeeAuth.DELETE("/delete", handler.Delete)
	}
}
