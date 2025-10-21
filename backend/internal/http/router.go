package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/http/routes"
)

func GlobalRouter(srv *handler.Server) *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r, srv)

	return r
}
