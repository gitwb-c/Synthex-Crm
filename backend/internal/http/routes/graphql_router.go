package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphqlRouter(r *gin.Engine, srv *handler.Server) {
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/playground", func(c *gin.Context) {
		playground.Handler("Playground", "/query").ServeHTTP(c.Writer, c.Request)
	})
}
