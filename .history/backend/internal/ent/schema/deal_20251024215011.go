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

type Deal struct {
	ent.Schema
}

func (Deal) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("title").NotEmpty().Unique().Annotations(entgql.QueryField(), entgql.OrderField("TITLE")),
		field.String("source").Annotations(entgql.QueryField()).Annotations(entgql.QueryField(), entgql.OrderField("SOURCE")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.QueryField(), entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.QueryField(), entgql.OrderField("UPDATED_AT")),
		field.UUID("tenantId", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
	}
}

func (Deal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Company.Type).Ref("deals").Field("tenantId").Unique(),
		edge.To("costumer", Costumer.Type).Unique(),
		edge.From("chat", Chat.Type).Ref("deal").Unique(),
		edge.To("stage", Stage.Type).Required().Unique().Required(),
		edge.From("dealCrmFields", DealCrmField.Type).Ref("deal"),
	}
}

func (Deal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (Deal) Mixins() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}

func  (Deal) Policyas