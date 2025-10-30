package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/bootstrap"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/contracts"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/services"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

func bootstrapRouter(r *gin.Engine, client *ent.Client, companyservice *services.CompanyService, employeeauthservice *services.EmployeeAuthService, employeeservice *services.EmployeeService, departmentservice *services.DepartmentService) {
	r.POST("/bootstrap", func(ctx *gin.Context) {
		var req contracts.BootstrapDtoRequest

		var err error
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
			return
		}
		setup, err := bootstrap.Bootstrap(client, companyservice, employeeauthservice, employeeservice, departmentservice, req, ctx.Request.Context())
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("setup completed successfully: %v", setup.Time)})
	})
}
