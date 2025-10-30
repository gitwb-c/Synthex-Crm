package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Enum("sentBy").Values("employee", "costumer").Annotations(entgql.QueryField(), entgql.OrderField("SENT_BY")),
		field.Bool("private").Annotations(entgql.QueryField()).Annotations(entgql.OrderField("PRIVATE")),
		field.Enum("type").Values("text", "audio", "image", "file").Annotations(entgql.QueryField(), entgql.OrderField("TYPE")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chat", Chat.Type).Unique(),
		edge.To("employee", Employee.Type),
		edge.From("text", Text.Type).Ref("message").Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("file", File.Type).Ref("message").Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("tenant", Company.Type).Ref("messages").Field("tenantId").Unique().Required().Immutable(),
	}
}

func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
