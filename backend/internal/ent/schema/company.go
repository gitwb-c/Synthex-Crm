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

type Company struct {
	ent.Schema
}

func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("name").Unique().NotEmpty().Annotations(entgql.QueryField()).Annotations(entgql.OrderField("NAME")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employees", Employee.Type),
		edge.To("costumers", Costumer.Type),
		edge.To("deals", Deal.Type),
		edge.To("departments", Department.Type),
		edge.To("pipelines", Pipeline.Type),
		edge.To("crmFields", Pipeline.Type),
		edge.To("dealCrmFields", Pipeline.Type),
		edge.To("dropdownLists", Pipeline.Type),
		edge.To("employeeAuths", Pipeline.Type),
		edge.To("files", Pipeline.Type),
		edge.To("messages", Pipeline.Type),
		edge.To("queues", Pipeline.Type),
		edge.To("rbacs", Pipeline.Type),
		edge.To("stages", Pipeline.Type),
		edge.To("texts", Pipeline.Type),
	}
}

func (Company) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
