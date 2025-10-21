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

type Message struct {
	ent.Schema
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.Enum("sentBy").Values("employee", "costumer").Annotations(entgql.QueryField()),
		field.Bool("private").Annotations(entgql.QueryField()),
		field.Enum("type").Values("text", "audio", "image", "file").Annotations(entgql.QueryField()),
		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chat", Chat.Type).Unique(),
		edge.To("employee", Employee.Type),
		edge.From("text", Text.Type).Ref("message").Unique(),
		edge.From("file", File.Type).Ref("message").Unique(),
	}
}

func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
