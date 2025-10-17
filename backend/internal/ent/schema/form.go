package schema

import "entgo.io/ent"

// Form holds the schema definition for the Form entity.
type Form struct {
	ent.Schema
}

// Fields of the Form.
func (Form) Fields() []ent.Field {
	return nil
}

// Edges of the Form.
func (Form) Edges() []ent.Edge {
	return nil
}
