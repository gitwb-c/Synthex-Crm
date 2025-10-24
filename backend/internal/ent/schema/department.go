package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Department struct {
	ent.Schema
}

func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID"), entgql.QueryField()),
		field.String("name").NotEmpty().Unique().Annotations().Annotations(entgql.OrderField("NAME")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
		field.UUID("tenantId", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
	}
}

func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Company.Type).Ref("departments").Field("tenantId").Unique(),
		edge.From("employee", Employee.Type).Ref("department"),
		edge.From("queues", Queue.Type).Ref("department"),
		edge.From("rbacs", Rbac.Type).Ref("department"),
	}
}

func (Department) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (Department) Mixins() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
