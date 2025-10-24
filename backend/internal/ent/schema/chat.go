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
		field.UUID("tenantId", uuid.UUID{}).Annotations(entgql.Type("ID")).Optional(),
	}
}

func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("deal", Deal.Type).Unique(),
		edge.To("employees", Employee.Type),
		edge.From("messages", Message.Type).Ref("chat"),
		edge.From("tenant", Company.Type).Ref("chats").Field("tenantId").Unique(),
	}
}

func (Chat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (Chat) Mixins() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
