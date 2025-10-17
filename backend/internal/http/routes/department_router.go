package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func departmentRouter(r *gin.Engine) {
	department := r.Group("/department")

	departmentService := services.NewDepartmentService()
	handler := handlers.NewDepartmentHandler(departmentService)
	{
		department.GET("/read", handler.Read)
		department.POST("/create", handler.Create)
		department.PATCH("/update", handler.Update)
		department.DELETE("/delete", handler.Delete)
	}
}
