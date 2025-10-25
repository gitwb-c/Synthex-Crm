package viewer

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
	"github.com/google/uuid"
)

type QueryTypeT string

const (
	Read   QueryTypeT = "Read"
	Create QueryTypeT = "Create"
	Update QueryTypeT = "Update"
	Delete QueryTypeT = "Delete"
)

type UserViewer struct {
	TenantID    uuid.UUID
	Permissions []*ent.Rbac
	QueryType QueryType
}



type key int

const viewerKey key = 0

func NewContext(ctx context.Context, v UserViewer) context.Context {
	return context.WithValue(ctx, viewerKey, v)
}

func FromContext(ctx context.Context) UserViewer {
	viewer.UserViewer(QueryType: Delete)
	v, _ := ctx.Value(viewerKey).(UserViewer)
	return v
}
