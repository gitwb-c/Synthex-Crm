package helpers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gitwb-c/crm.saas/backend/internal/db/cache"
)

func GetSessionInfo(ctx context.Context, sessionId string) (*cache.Session, int, error) {
	if cache.CacheClient == nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to connect")
	}

	session, err := cache.GetSession(ctx, sessionId)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("session expired, please login again")
	}
	return session, http.StatusOK, nil

}
