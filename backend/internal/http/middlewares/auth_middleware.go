package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/http/middlewares/helpers"
	"github.com/gitwb-c/crm.saas/backend/pkg/jwtpkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authStr, err := ctx.Cookie("auth_jwt_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(authStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return jwtpkg.SecretKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			sessionId, exists := claims["sessionId"]
			if !exists {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
				ctx.Abort()
				return
			}
			session, status, err := helpers.GetSessionInfo(ctx.Request.Context(), sessionId.(string))
			if err != nil {
				ctx.JSON(status, gin.H{"error": err.Error()})
				ctx.Abort()
				return
			}
			ctx.Set("sessionId", sessionId.(string))
			ctx.Set("employeeId", session.EmployeeId)
			ctx.Set("tenantId", session.TenantId)
			ctx.Set("departmentId", session.DepartmentId)

		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
