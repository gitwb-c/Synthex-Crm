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

type Stage struct {
	ent.Schema
}

func (Stage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID"), entgql.QueryField()),
		field.String("name").NotEmpty().Annotations(entgql.QueryField()),
		field.String("color").Default("#ffffff").NotEmpty().Annotations(),
		field.Bool("lossOrGain").Annotations(entgql.QueryField()),
		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Stage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pipeline", Pipeline.Type).Required().Unique(),
		edge.From("deals", Deal.Type).Ref("stage"),
		edge.To("queue", Queue.Type).Unique(),
	}
}

func (Stage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
