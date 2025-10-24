package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/http/middlewares"
)

func graphqlRouter(r *gin.Engine, srv *handler.Server) {
	r.POST("/query/:tenantId", middlewares.AuthMiddleware(), middlewares.TenantMiddleware(), func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	})

	r.POST("/query", middlewares.AuthMiddleware(), middlewares.TenantMiddleware(), func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	})

	r.GET("/playground", func(ctx *gin.Context) {
		playground.Handler("Playground", "/query/f194757e-1b67-4877-985e-f4a4e3311eec").ServeHTTP(ctx.Writer, ctx.Request)
	})
}
