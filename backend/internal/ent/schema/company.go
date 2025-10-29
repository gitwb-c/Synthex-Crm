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

type Company struct {
	ent.Schema
}

func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.String("name").Unique().NotEmpty().Annotations(entgql.QueryField()).Annotations(entgql.OrderField("NAME")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employees", Employee.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("costumers", Costumer.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("deals", Deal.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("chats", Chat.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("departments", Department.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("pipelines", Pipeline.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("crmFields", CrmField.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("dealCrmFields", DealCrmField.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("dropdownLists", DropdownList.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("employeeAuths", EmployeeAuth.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("files", File.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("messages", Message.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("queues", Queue.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("rbacs", Rbac.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("stages", Stage.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("texts", Text.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Company) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
