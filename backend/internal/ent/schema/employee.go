package schema

import "entgo.io/ent"

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return nil
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return nil
}
