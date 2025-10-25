package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TenantMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tenantIDStr, exists := ctx.Get("tenantId")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing tenant"})
			ctx.Abort()
			return
		}
		tenantID, err := uuid.Parse(tenantIDStr.(string))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenant"})
			ctx.Abort()
			return
		}
		if ctx.Param("tenantId") != tenantID.String() {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "mismatched tenant"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
