package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/pkg/jwtpkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authStr, err := ctx.Cookie("auth")
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

		// if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		// 	ctx.Set("employeId", claims["employeeId"])
		// 	ctx.Set("companyId", claims["companyId"])
		// 	ctx.Set("departmentId", claims["departmentId"])
		// }
		ctx.Next()
	}
}
