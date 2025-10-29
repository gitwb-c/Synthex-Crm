package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/wire"
)

func RegisterRoutes(r *gin.Engine, client *ent.Client, srv *handler.Server) {
	companyservice := wire.InitializeCompanyService(client)
	departmentervice := wire.InitializeDepartmentService(client)
	employeeauthservice := wire.InitializeEmployeeAuthService(client)
	employeeservice := wire.InitializeEmployeeService(client)

	graphqlRouter(r, srv, employeeservice, departmentervice)
	loginRouter(r, employeeauthservice)
	bootstrapRouter(r, client, companyservice, employeeauthservice, employeeservice, departmentervice)
}
