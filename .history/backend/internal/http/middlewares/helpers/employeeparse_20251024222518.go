package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/services"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

func EmployeeParse(ctx *gin.Context, employeeservice *services.EmployeeService) (int, *ent.Employee, map[string]any) {
	employeeIdStr, exists := ctx.Get("employeeId")
	if !exists {
		return http.StatusUnauthorized, nil, gin.H{"error": "employeeId not found"}
	}
	employeeId, err := uuid.Parse(employeeIdStr.(string))
	if err != nil {
		return http.StatusBadRequest, nil, gin.H{"error": "invalid employeeId"}
	}
	employee, err := employeeservice.ReadID(ctx, employeeId)
	if err != nil {
		return http.StatusBadRequest, nil, gin.H{"error": err.Error()}
	}
	if employee == nil {
		return http.StatusBadRequest, nil, gin.H{"error": "employee not found"}
	}
	if employee.EmploymentStatus.String() != "active" {
		return http.StatusBadRequest, nil, gin.H{"error": "employee is not active"}
	}
	return http.StatusOK, employee, nil
}
