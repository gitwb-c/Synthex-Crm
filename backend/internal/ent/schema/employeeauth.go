package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type EmployeeAuth struct {
	ent.Schema
}

func (EmployeeAuth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("name").NotEmpty().Annotations(entgql.Skip()),
		field.String("email").NotEmpty().Unique().Sensitive().Annotations(entgql.Skip()),
		field.String("password").NotEmpty().Sensitive().Annotations(entgql.Skip()),
		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (EmployeeAuth) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (EmployeeAuth) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
