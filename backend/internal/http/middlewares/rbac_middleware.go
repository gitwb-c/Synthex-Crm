package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/services"
	"github.com/gitwb-c/crm.saas/backend/internal/graphql/graph/graph_helpers"
	"github.com/gitwb-c/crm.saas/backend/internal/http/middlewares/helpers"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
)

func RbacMiddleware(employeeservice *services.EmployeeService, departmentservice *services.DepartmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employeestatus, employee, er := helpers.EmployeeParse(ctx, employeeservice)
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

		querytypestatus, querytype, err := graph_helpers.GetQueryTypeCtx(ctx)
		if querytypestatus != http.StatusOK {
			ctx.JSON(querytypestatus, err)
			ctx.Abort()
			return
		}

		reqCtx := viewer.NewContext(ctx.Request.Context(), viewer.UserViewer{TenantID: employee.TenantId, Permissions: department.Edges.Rbacs, QueryType: viewer.QueryType(querytype)})
		ctx.Request = ctx.Request.WithContext(reqCtx)

		ctx.Next()
	}
}
