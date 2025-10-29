package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type CrmField struct {
	ent.Schema
}

func (CrmField) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("name").NotEmpty().Annotations(entgql.QueryField(), entgql.OrderField("NAME")),
		field.String("section").Annotations(entgql.QueryField(), entgql.OrderField("SECTION")),
		field.Enum("type").Values("txt", "date", "checkbox", "dropdownList").Immutable().Annotations(entgql.QueryField(), entgql.OrderField("TYPE")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (CrmField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dropdownList", DropdownList.Type).Ref("crmField").Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("dealCrmField", DealCrmField.Type).Ref("crmField").Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("tenant", Company.Type).Ref("crmFields").Field("tenantId").Unique().Required().Immutable(),
	}
}

func (CrmField) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func CrmIndexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("tenant").Unique(),
	}
}

func (CrmField) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
