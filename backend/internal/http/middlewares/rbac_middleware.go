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

		permissions := []string{}
		for _, rbac := range department.Edges.Rbacs {
			permissions = append(permissions, string(rbac.Access))
		}

		sessionId, exists := ctx.Get("sessionId")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing session"})
			ctx.Abort()
			return
		}

		reqCtx := viewer.NewContext(ctx.Request.Context(), viewer.UserViewer{
			TenantID:    employee.TenantId,
			SessionID:   sessionId.(string),
			Permissions: &permissions,
		})

		ctx.Request = ctx.Request.WithContext(reqCtx)
		ctx.Next()
	}
}
