package schema

import "entgo.io/ent"

// Pipeline holds the schema definition for the Pipeline entity.
type Pipeline struct {
	ent.Schema
}

// Fields of the Pipeline.
func (Pipeline) Fields() []ent.Field {
	return nil
}

// Edges of the Pipeline.
func (Pipeline) Edges() []ent.Edge {
	return nil
}
