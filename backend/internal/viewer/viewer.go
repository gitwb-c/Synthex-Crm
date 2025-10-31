package viewer

import (
	"context"

	"github.com/google/uuid"
)

type QueryType string

const (
	Read   QueryType = "Read"
	Create QueryType = "Create"
	Update QueryType = "Update"
	Delete QueryType = "Delete"
)

type Signature string

const (
	Bootstrap Signature = "Bootstrap"
	Login     Signature = "Login"
	Common    Signature = "Common"
)

type UserViewer struct {
	TenantID    uuid.UUID
	SessionID   string
	Permissions *[]string
	Signature   *Signature
}

type key int

const ViewerKey key = 0

func NewContext(ctx context.Context, v UserViewer) context.Context {
	return context.WithValue(ctx, ViewerKey, v)
}

func FromContext(ctx context.Context) UserViewer {
	v, _ := ctx.Value(ViewerKey).(UserViewer)
	return v
}
