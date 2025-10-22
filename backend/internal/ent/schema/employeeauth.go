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

type EmployeeAuth struct {
	ent.Schema
}

func (EmployeeAuth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty().Unique().Sensitive(),
		field.String("password").NotEmpty().Sensitive(),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (EmployeeAuth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("employee", Employee.Type).Ref("employeeAuth").Unique(),
	}
}

func (EmployeeAuth) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
