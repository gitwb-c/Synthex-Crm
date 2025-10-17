package schema

import "entgo.io/ent"

// Stage holds the schema definition for the Stage entity.
type Stage struct {
	ent.Schema
}

// Fields of the Stage.
func (Stage) Fields() []ent.Field {
	return nil
}

// Edges of the Stage.
func (Stage) Edges() []ent.Edge {
	return nil
}
