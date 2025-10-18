package routes

import (
	"crm.saas/backend/internal"
	"github.com/gin-gonic/gin"
)

func departmentRouter(r *gin.Engine) {
	department := r.Group("/department")

	handler := internal.InitializeDepartmentHandler()
	{
		department.GET("/", handler.Read)
		department.POST("/", handler.Create)
		department.PATCH("/", handler.Update)
		department.DELETE("/", handler.Delete)
	}
}
