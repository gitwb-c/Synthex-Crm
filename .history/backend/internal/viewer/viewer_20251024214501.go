package viewer

import (
	"context"

	"github.com/google/uuid"
)

type UserViewer struct {
	TenantID uuid.UUID
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
