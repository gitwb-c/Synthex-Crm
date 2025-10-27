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
		edge.To("employees", Employee.Type),
		edge.To("costumers", Costumer.Type),
		edge.To("deals", Deal.Type),
		edge.To("chats", Chat.Type),
		edge.To("departments", Department.Type),
		edge.To("pipelines", Pipeline.Type),
		edge.To("crmFields", CrmField.Type),
		edge.To("dealCrmFields", DealCrmField.Type),
		edge.To("dropdownLists", DropdownList.Type),
		edge.To("employeeAuths", EmployeeAuth.Type),
		edge.To("files", File.Type),
		edge.To("messages", Message.Type),
		edge.To("queues", Queue.Type),
		edge.To("rbacs", Rbac.Type),
		edge.To("stages", Stage.Type),
		edge.To("texts", Text.Type),
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

// func (Company) Policy() ent.Policy {
// 	permissions := map[viewer.QueryType][]string{
// 		viewer.Read:   {"manage_system_settings"},
// 		viewer.Create: {"manage_system_settings"},
// 		viewer.Update: {"manage_system_settings"},
// 		viewer.Delete: {"manage_system_settings"},
// 	}
// 	return privacy.Policy{
// 		Mutation: privacy.MutationPolicy{
// 			rule.MutationRules(permissions),
// 		},
// 		Query: privacy.QueryPolicy{
// 			rule.QueryRules(permissions),
// 		},
// 	}
// }
