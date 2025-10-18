package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func employeeAuthRouter(r *gin.Engine) {
	employeeAuth := r.Group("/employeeauth")

	handler := internal.InitializeEmployeeAuthHandler()
	{
		employeeAuth.GET("/", handler.Read)
		employeeAuth.POST("/", handler.Create)
		employeeAuth.PATCH("/", handler.Update)
		employeeAuth.DELETE("/", handler.Delete)
	}
}
