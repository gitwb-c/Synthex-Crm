package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Pipeline struct {
	ent.Schema
}

func (Pipeline) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID"), entgql.QueryField()),
		field.String("name").NotEmpty().Annotations(entgql.QueryField()).Annotations(entgql.OrderField("NAME")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (Pipeline) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Company.Type).Ref("pipelines").Field("tenantId").Unique().Required().Immutable(),
		edge.From("stages", Stage.Type).Ref("pipeline").Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Pipeline) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func PipelineIndexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("tenant").Unique(),
	}
}

func (Pipeline) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
