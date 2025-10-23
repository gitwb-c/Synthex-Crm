package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/schema/mixin"
	"github.com/google/uuid"
)

type Text struct {
	ent.Schema
}

func (Text) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("text").NotEmpty().Annotations(entgql.QueryField(), entgql.OrderField("TEXT")),
	}
}

func (Text) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("message", Message.Type).Unique().Required(),
		edge.From("company", Company.Type).Ref("texts").Field("tenant_id").Required().Unique(),
	}
}

func (Text) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Text) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixin.TenantMixin{},
	}
}
