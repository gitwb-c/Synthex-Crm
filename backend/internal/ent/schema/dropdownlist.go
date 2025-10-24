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

type DropdownList struct {
	ent.Schema
}

func (DropdownList) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("value").NotEmpty().Annotations(entgql.QueryField()).Annotations(entgql.OrderField("VALUE")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
		field.UUID("tenantId", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
	}
}

func (DropdownList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("crmField", CrmField.Type).Required(),
		edge.From("tenant", Company.Type).Ref("dropdownLists").Field("tenantId").Unique(),
	}
}

func (DropdownList) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (DropdownList) Mixins() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
