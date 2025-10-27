package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/wire"
)

func RegisterRoutes(r *gin.Engine, client *ent.Client, srv *handler.Server) {
	employeeservice := wire.InitializeEmployeeService(client)
	departmentervice := wire.InitializeDepartmentService(client)
	graphqlRouter(r, srv, employeeservice, departmentervice)

	employeeauthservice := wire.InitializeEmployeeAuthService(client)
	loginRouter(r, employeeauthservice)
}
