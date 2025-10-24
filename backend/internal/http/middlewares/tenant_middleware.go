package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
	"github.com/google/uuid"
)

func TenantMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tenantIDStr, exists := ctx.Get("tenantId")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing tenantId"})
			ctx.Abort()
			return
		}
		tenantID, err := uuid.Parse(tenantIDStr.(string))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenantId"})
			ctx.Abort()
			return
		}
		if ctx.Param("tenantId") != tenantID.String() {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenantId"})
			ctx.Abort()
			return
		}
		reqCtx := viewer.NewContext(ctx.Request.Context(), viewer.UserViewer{TenantID: tenantID})
		ctx.Request = ctx.Request.WithContext(reqCtx)
		ctx.Next()
	}
}
