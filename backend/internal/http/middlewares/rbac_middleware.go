package middlewares

import "github.com/gin-gonic/gin"

func RbacMiddleware(requiredPermissions []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}
