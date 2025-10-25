package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/services"
	"github.com/gitwb-c/crm.saas/backend/internal/http/middlewares/helpers"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
)

func RbacMiddleware(employeeservice *services.EmployeeService, departmentservice *services.DepartmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employeestatus, er := helpers.EmployeeParse(ctx, employeeservice)
		if employeestatus != http.StatusOK {
			ctx.JSON(employeestatus, er)
			ctx.Abort()
			return
		}

		departmentstatus, department, err := helpers.DepartmentParse(ctx, departmentservice)
		if departmentstatus != http.StatusOK {
			ctx.JSON(departmentstatus, err)
			ctx.Abort()
			return
		}

		ctx.Set("permissions", department.Edges.Rbacs)
		reqCtx := viewer.NewContext(ctx.Request.Context(), viewer.UserViewer{Permissions: department.Edges.Rbacs})

		ctx.Next()
	}
}
