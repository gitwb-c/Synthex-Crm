package http

import (
	"crm.saas/backend/internal/http/routes"
	"github.com/gin-gonic/gin"
)

func GlobalRouter() *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r)

	return r
}
