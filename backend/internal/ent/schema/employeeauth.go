package schema

import "entgo.io/ent"

// EmployeeAuth holds the schema definition for the EmployeeAuth entity.
type EmployeeAuth struct {
	ent.Schema
}

// Fields of the EmployeeAuth.
func (EmployeeAuth) Fields() []ent.Field {
	return nil
}

// Edges of the EmployeeAuth.
func (EmployeeAuth) Edges() []ent.Edge {
	return nil
}
