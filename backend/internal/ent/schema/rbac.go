package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/schema/mixin"
	"github.com/google/uuid"
)

type Rbac struct {
	ent.Schema
}

func (Rbac) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Annotations(entgql.Type("ID")),
		field.Enum("access").
			Values(
				// Allows viewing all deals without restriction
				"view_all_deals",

				// Allows viewing only deals that the employee is responsible for in the chat
				"view_responsible_chat_deals",

				// Allows viewing only deals that are in the department's phase
				"view_department_phase_deals",

				// Allows deleting any card from the system
				"delete_deals",

				// Allows changing the responsible person for the chat or business of a card
				"change_responsible_chat_or_business",

				// Allows viewing all customer data, including sensitive and private information
				"view_all_customer_data",

				// Allows viewing customer data specific to the user's department
				"view_department_customer_data",

				// Allows creating new deals in the system (e.g., creating sales opportunities or tasks)
				"create_deal",

				// Allows editing existing deals (e.g., updating customer details or business phase)
				"edit_deals",

				// Allows approving or denying deals (e.g., for validation or workflow approval)
				"approve_or_deny_deals",

				// Allows viewing the history of interactions with customers, such as calls, emails, or meetings
				"view_chat_history",

				// Allows viewing advanced or custom reports (e.g., sales performance or financial reports)
				"view_advanced_reports",

				// Allows creating, editing, or deleting users in the CRM system (admin-level permission)
				"manage_users",

				// Allows managing departments, including adding or removing departments and editing their information
				"manage_departments",

				// Allows modifying system-wide settings and preferences (admin-level permission)
				"manage_system_settings",

				// Allows managing products or services offered within the CRM (e.g., adding or updating product listings)
				"manage_products_or_services",

				// Allows exporting data from the CRM (e.g., customer data, deals, reports)
				"export_data",

				// Allows configuring and managing workflows within the CRM (e.g., sales processes, task automation)
				"configure_workflows",

				// Allows sending notifications to users or customers within the system (e.g., alerts, updates)
				"send_notifications",

				// Allows viewing system notifications and alerts
				"view_notifications",

				// Allows configuring integrations with external systems or services (e.g., email, payment gateways)
				"configure_integrations",

				// Allows accessing business analytics data to identify trends, performance, and insights
				"view_dashboard_analytics_data",

				// Allows performing post-sale actions, such as follow-up, re-open deals, or close the business cycle
				"post_sale_actions",
			),
		field.Time("createdAt").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (Rbac) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("department", Department.Type).Required().Unique(),
		edge.From("company", Company.Type).Ref("rbacs").Field("tenant_id").Required().Unique(),
	}
}

func (Rbac) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (Rbac) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixin.TenantMixin{},
	}
}
