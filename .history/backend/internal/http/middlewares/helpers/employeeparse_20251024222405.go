package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/services"
	"github.com/google/uuid"
)

func EmployeeParse(ctx *gin.Context, employeeservice *services.EmployeeService) (int, employee, map[string]any) {
	employeeIdStr, exists := ctx.Get("employeeId")
	if !exists {
		return http.StatusUnauthorized, gin.H{"error": "employeeId not found"}
	}
	employeeId, err := uuid.Parse(employeeIdStr.(string))
	if err != nil {
		return http.StatusBadRequest, gin.H{"error": "invalid employeeId"}
	}
	employee, err := employeeservice.ReadID(ctx, employeeId)
	if err != nil {
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	}
	if employee == nil {
		return http.StatusBadRequest, gin.H{"error": "employee not found"}
	}
	if employee.EmploymentStatus.String() != "active" {
		return http.StatusBadRequest, gin.H{"error": "employee is not active"}
	}
	return http.StatusOK, employee, nil
}
