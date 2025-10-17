package routes

import (
	"crm.saas/backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func departmentRouter(r *gin.Engine) {
	department := r.Group("/department")
	handler := handlers.NewDepartmentHandler()

	{
		department.GET("/read", handler.Read)
		department.POST("/create", handler.Create)
		department.PATCH("/update", handler.Update)
		department.DELETE("/delete", handler.Delete)
	}
}
