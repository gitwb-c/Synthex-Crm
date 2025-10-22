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

type Queue struct {
	ent.Schema
}

func (Queue) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID"), entgql.QueryField()),
		field.String("name").NotEmpty().Unique().Annotations(entgql.QueryField(), entgql.OrderField("NAME")),
		field.Enum("type").Values("ring").Default("ring").Annotations(entgql.QueryField(), entgql.OrderField("TYPE")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (Queue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stages", Stage.Type).Ref("queue"),
		edge.From("employees", Employee.Type).Ref("queues"),
		edge.To("department", Department.Type).Required(),
	}
}

func (Queue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
