package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Employee struct {
	ent.Schema
}

func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID"), entgql.QueryField()),
		field.String("name").NotEmpty().Annotations(entgql.QueryField(), entgql.OrderField("NAME")),
		field.Enum("employmentStatus").Values("active", "terminated", "onLeave").Annotations(entgql.QueryField()).Annotations(entgql.OrderField("EMPLOYMENT_STATUS")),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employeeAuth", EmployeeAuth.Type).Unique().Required(),
		edge.From("tenant", Company.Type).Ref("employees").Field("tenantId").Unique().Required().Immutable(),
		edge.To("department", Department.Type).Required().Unique(),
		edge.From("chat", Chat.Type).Ref("employees"),
		edge.To("queues", Queue.Type),
		edge.From("messages", Message.Type).Ref("employee"),
	}
}

func (Employee) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func EmployeeIndexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("tenant").Unique(),
	}
}

func (Employee) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}
