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

type CrmField struct {
	ent.Schema
}

func (CrmField) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("name").NotEmpty().Unique().Annotations(entgql.QueryField()),
		field.String("section").Annotations(entgql.QueryField()),
		field.Enum("type").Values("txt", "date", "checkbox", "dropdownList").Immutable().Annotations(entgql.QueryField()),
		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (CrmField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dropdownList", DropdownList.Type).Ref("crmField"),
		edge.From("dealCrmField", DealCrmField.Type).Ref("crmField"),
	}
}

func (CrmField) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
