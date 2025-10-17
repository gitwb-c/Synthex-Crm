package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func employeeRouter(r *gin.Engine) {
	employee := r.Group("/employee")

	employeeService := services.NewEmployeeService()
	handler := handlers.NewEmployeeHandler(employeeService)
	{
		employee.GET("/read", handler.Read)
		employee.POST("/create", handler.Create)
		employee.PATCH("/update", handler.Update)
		employee.DELETE("/delete", handler.Delete)
	}
}
