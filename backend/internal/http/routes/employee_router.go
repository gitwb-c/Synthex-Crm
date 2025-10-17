package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func employeeRouter(r *gin.Engine) {
	employee := r.Group("/employee")
	handler := handlers.NewEmployeeHandler()

	{
		employee.GET("/read", handler.Read)
		employee.POST("/create", handler.Create)
		employee.PATCH("/update", handler.Update)
		employee.DELETE("/delete", handler.Delete)
	}
}
