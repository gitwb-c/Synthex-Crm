package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type File struct {
	ent.Schema
}

func (File) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("url").NotEmpty().Annotations(entgql.QueryField()).Annotations(entgql.OrderField("URL")),
		field.String("caption").Optional().Annotations(entgql.QueryField()),
		field.String("mimeType").Annotations(entgql.QueryField()),
		field.String("fileName").Annotations(entgql.QueryField()),
	}
}

func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("message", Message.Type).Unique().Required(),
		edge.From("tenant", Company.Type).Ref("files").Field("tenantId").Unique().Required().Immutable().Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
