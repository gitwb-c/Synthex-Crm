package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func employeeAuthRouter(r *gin.Engine) {
	employeeAuth := r.Group("/employeeauth")

	employeeService := services.NewEmployeeAuthService()
	handler := handlers.NewEmployeeAuthHandler(employeeService)

	{
		employeeAuth.GET("/read", handler.Read)
		employeeAuth.POST("/create", handler.Create)
		employeeAuth.PATCH("/update", handler.Update)
		employeeAuth.DELETE("/delete", handler.Delete)
	}
}
