package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/schema/mixin"
	"github.com/google/uuid"
)

type Costumer struct {
	ent.Schema
}

func (Costumer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("name").NotEmpty().Annotations().Annotations(entgql.OrderField("NAME")),
		field.String("phone").NotEmpty().Annotations(entgql.QueryField()).Unique().Annotations(entgql.OrderField("PHONE")),
		field.String("email").NotEmpty().Annotations(entgql.QueryField()).Unique().Annotations(entgql.OrderField("EMAIL")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (Costumer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).Ref("costumers").Field("tenant_id").Required().Unique(),
		edge.From("deals", Deal.Type).Ref("costumer"),
	}
}

func (Costumer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (Costumer) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixin.TenantMixin{},
	}
}
