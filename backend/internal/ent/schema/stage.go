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
		field.String("name").NotEmpty().Annotations(entgql.QueryField(), entgql.OrderField("NAME")),
		field.String("color").Default("#ffffff").NotEmpty(),
		field.Bool("lossOrGain").Annotations(entgql.QueryField(), entgql.OrderField("LOSS_OR_GAIN")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (Stage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pipeline", Pipeline.Type).Required().Unique(),
		edge.From("deals", Deal.Type).Ref("stage"),
		edge.To("queue", Queue.Type).Unique(),
		edge.From("tenant", Company.Type).Ref("stages").Field("tenantId").Unique().Required().Immutable(),
	}
}

func (Stage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (Stage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
