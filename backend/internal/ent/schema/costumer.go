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

type Costumer struct {
	ent.Schema
}

func (Costumer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("name").NotEmpty().Annotations(),
		field.String("phone").NotEmpty().Annotations(entgql.QueryField()),
		field.String("email").NotEmpty().Annotations(entgql.QueryField()),
		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").UpdateDefault(time.Now),
	}
}

func (Costumer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deals", Deal.Type).Ref("costumer"),
	}
}

func (Costumer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
