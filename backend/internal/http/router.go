package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/http/routes"
)

func GlobalRouter(client *ent.Client, srv *handler.Server) *gin.Engine {
	r := gin.Default()

	routes.RegisterRoutes(r, client, srv)

	return r
}
