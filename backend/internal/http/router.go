package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/http/routes"
	"github.com/redis/go-redis/v9"
)

func GlobalRouter(client *ent.Client, cacheClient *redis.Client, srv *handler.Server) *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r, client, cacheClient, srv)

	return r
}
