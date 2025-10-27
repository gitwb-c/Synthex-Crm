package viewer

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

type QueryType string

const (
	Read   QueryType = "Read"
	Create QueryType = "Create"
	Update QueryType = "Update"
	Delete QueryType = "Delete"
)

type UserViewer struct {
	TenantID    uuid.UUID
	Permissions []*ent.Rbac
}

type key int

const viewerKey key = 0

func NewContext(ctx context.Context, v UserViewer) context.Context {
	return context.WithValue(ctx, viewerKey, v)
}

func FromContext(ctx context.Context) UserViewer {
	v, _ := ctx.Value(viewerKey).(UserViewer)
	return v
}
