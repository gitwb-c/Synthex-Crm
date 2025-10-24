package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/services"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

func DepartmentParse(ctx *gin.Context, departmentservice *services.DepartmentService) (int, *ent.Department, map[string]any) {
	departmentIdStr, exists := ctx.Get("departmentId")
	if !exists {
		return http.StatusUnauthorized, nil, gin.H{"error": "departmentId not found"}
	}
	departmentId, err := uuid.Parse(departmentIdStr.(string))
	if err != nil {
		return http.StatusBadRequest, nil, gin.H{"error": "invalid departmentId"}
	}
	department, err := departmentservice.ReadRbacs(ctx, departmentId)
	if err != nil {
		return http.StatusBadRequest, nil, gin.H{"error": err.Error()}
	}
	if department == nil {
		return http.StatusBadRequest, nil, gin.H{"error": "department not found"}
	}
	return http.StatusOK, department, nil
}
