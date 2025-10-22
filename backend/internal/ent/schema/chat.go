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

type Chat struct {
	ent.Schema
}

func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("title").NotEmpty().Annotations(entgql.OrderField("TITLE")),
		field.Bool("accepted").Annotations(entgql.OrderField("ACCEPTED")),
		field.Bool("locked").Annotations(entgql.OrderField("LOCKED")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("deal", Deal.Type).Unique(),
		edge.To("employees", Employee.Type),
		edge.From("messages", Message.Type).Ref("chat"),
	}
}

func (Chat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
