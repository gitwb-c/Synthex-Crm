package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gitwb-c/crm.saas/backend/internal/db/cache"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/contracts"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/services"
	"github.com/google/uuid"
)

func loginRouter(r *gin.Engine, service *services.EmployeeAuthService) {
	auth := r.Group("/auth")

	auth.POST("/login/:tenantId", func(ctx *gin.Context) {
		tenantIdStr := ctx.Param("tenantId")
		tenantId, err := uuid.Parse(tenantIdStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req contracts.LoginDtoRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newLogin, status, err := service.ValidateLogin(ctx, req.Email, req.Password, tenantId)
		if err != nil {
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}
		if err := cache.SaveSession(ctx.Request.Context(), newLogin.SessionId, cache.Session{
			EmployeeId:   newLogin.EmployeeId,
			TenantId:     newLogin.TenantId,
			DepartmentId: newLogin.DepartmentId,
			IP:           ctx.ClientIP(),
			ExpiresAt:    time.Now().Add(7 * 24 * time.Hour)}); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.SetCookie("auth_jwt_token", newLogin.Jwt, 1814400, "/", os.Getenv("WEB_DOMAIN"), false, true)
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("logged in at: %v", newLogin.Time.Format("2006-01-02 15:04:05"))})
	})
}
