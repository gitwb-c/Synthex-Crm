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

type DealCrmField struct {
	ent.Schema
}

func (DealCrmField) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("value").NotEmpty().Annotations(entgql.QueryField()).Annotations(entgql.OrderField("VALUE")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
		field.UUID("tenantId", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
	}
}

func (DealCrmField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("deal", Deal.Type).Unique().Required(),
		edge.To("crmField", CrmField.Type).Unique().Required(),
		edge.From("tenant", Company.Type).Ref("dealCrmFields").Field("tenantId").Unique(),
	}
}

func (DealCrmField) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (DealCrmField) Mixins() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
