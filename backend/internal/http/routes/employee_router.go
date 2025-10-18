package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func employeeRouter(r *gin.Engine) {
	employee := r.Group("/employee")

	handler := internal.InitializeEmployeeHandler()
	{
		employee.GET("/", handler.Read)
		employee.POST("/", handler.Create)
		employee.PATCH("/", handler.Update)
		employee.DELETE("/", handler.Delete)
	}
}
