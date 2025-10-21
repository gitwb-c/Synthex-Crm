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

type Deal struct {
	ent.Schema
}

func (Deal) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("title").NotEmpty().Unique().Annotations(entgql.QueryField()),
		field.String("source").Annotations(entgql.QueryField()),
		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Deal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("costumer", Costumer.Type).Required().Unique(),
		edge.From("chat", Chat.Type).Ref("deal").Unique(),
		edge.To("stage", Stage.Type).Required().Unique(),
		edge.From("dealCrmFields", DealCrmField.Type).Ref("deal"),
	}
}

func (Deal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
