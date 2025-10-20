package http

import (
	"crm.saas/backend/internal/http/routes"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func GlobalRouter(srv *handler.Server) *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r, srv)

	return r
}
