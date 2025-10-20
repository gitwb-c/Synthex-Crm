package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, srv *handler.Server) {
	graphqlRouter(r, srv)
}
