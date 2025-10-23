package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphqlRouter(r *gin.Engine, srv *handler.Server) {
	r.POST("/query", func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	})

	r.GET("/playground", func(ctx *gin.Context) {
		playground.Handler("Playground", "/query").ServeHTTP(ctx.Writer, ctx.Request)
	})
}
