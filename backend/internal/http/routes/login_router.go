package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/contracts"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/services"
	"github.com/google/uuid"
)

func loginRouter(r *gin.Engine, service *services.EmployeeAuthService) {
	r.POST("/auth/login/:tenantId", func(ctx *gin.Context) {
		var req contracts.LoginDtoRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tenantIdStr := ctx.Param("tenantId")
		tenantId, err := uuid.Parse(tenantIdStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newLogin, err := service.ValidateLogin(ctx, req.Email, req.Password, tenantId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.SetCookie("auth_jwt_token", newLogin.Jwt, 1814400, "/", os.Getenv("WEB_DOMAIN"), false, true)
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("logged in at: %v", newLogin.Time.Format("2006-01-02 15:04:05"))})
	})
}
