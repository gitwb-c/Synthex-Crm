package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/services"
	"github.com/gitwb-c/crm.saas/backend/internal/http/middlewares"
)

func graphqlRouter(r *gin.Engine, srv *handler.Server, employeeservice *services.EmployeeService, departmentservice *services.DepartmentService) {
	graphql := r.Group("/graphql")

	graphql.POST("/query/:tenantId", middlewares.AuthMiddleware(), middlewares.TenantMiddleware(), middlewares.RbacMiddleware(employeeservice, departmentservice), func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	})

	graphql.GET("/playground", func(ctx *gin.Context) {
		playground.Handler("Playground", "/graphql/query/30e1f1a3-fb3b-4de7-a875-2db37b8e96d3").ServeHTTP(ctx.Writer, ctx.Request)
	})
}
