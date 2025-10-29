package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
	}
}

func (Deal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Company.Type).Ref("deals").Field("tenantId").Unique().Required().Immutable(),
		edge.To("costumer", Costumer.Type).Unique(),
		edge.From("chat", Chat.Type).Ref("deal").Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("stage", Stage.Type).Required().Unique().Required(),
		edge.From("dealCrmFields", DealCrmField.Type).Ref("deal").Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
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

func (Deal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
