package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Deal struct {
	ent.Schema
}

func (Deal) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Annotations(entgql.OrderField("TITLE")),
	}
}

func (Deal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
