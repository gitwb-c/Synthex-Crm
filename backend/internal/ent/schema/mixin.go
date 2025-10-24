package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type TenantMixin struct {
	mixin.Schema
}

func (TenantMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("tenantId", uuid.UUID{}).Immutable(),
	}
}

func (TenantMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).
			Field("tenantId").
			Unique().
			Required().
			Immutable(),
	}
}
