package graph_helpers

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
	"github.com/google/uuid"
)

type key int

const graphkey key = 1

func AddQuery(ctx context.Context, querytype string) context.Context {
	return context.WithValue(ctx, graphkey, querytype)
}

func GetQuery(ctx context.Context) (bool, string) {
	query := ctx.Value(graphkey)
	if query == nil {
		return false, ""

	}
	querytypeStr, ok := query.(string)
	if !ok {
		return false, ""
	}
	return true, querytypeStr
}

func GetViewer(ctx context.Context) (bool, *viewer.UserViewer) {
	v, ok := ctx.Value(viewer.ViewerKey).(viewer.UserViewer)
	if !ok {
		return false, nil
	}

	if v.TenantID == uuid.Nil {
		return false, nil
	}

	if v.SessionID == "" {
		return false, nil
	}

	return true, &v
}
