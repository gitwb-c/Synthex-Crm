package viewer

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

// QueryTypeT define o tipo da operação que o viewer está executando.
type QueryTypeT string

const (
	Read   QueryTypeT = "Read"
	Create QueryTypeT = "Create"
	Update QueryTypeT = "Update"
	Delete QueryTypeT = "Delete"
)

// UserViewer representa o contexto do usuário atual.
type UserViewer struct {
	TenantID    uuid.UUID
	Permissions []*ent.Rbac
	QueryType   QueryTypeT
}

// chave privada para evitar colisão no context
type key struct{}

var viewerKey key

// NewContext cria um novo contexto com o viewer associado.
func NewContext(ctx context.Context, v UserViewer) context.Context {
	return context.WithValue(ctx, viewerKey, v)
}

// FromContext retorna o viewer armazenado no contexto.
// Se não houver viewer, retorna um valor vazio.
func FromContext(ctx context.Context) UserViewer {
	v, ok := ctx.Value(viewerKey).(UserViewer)
	if !ok {
		return UserViewer{}
	}
	return v
}
