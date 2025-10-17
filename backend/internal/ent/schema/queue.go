package schema

import "entgo.io/ent"

// Queue holds the schema definition for the Queue entity.
type Queue struct {
	ent.Schema
}

// Fields of the Queue.
func (Queue) Fields() []ent.Field {
	return nil
}

// Edges of the Queue.
func (Queue) Edges() []ent.Edge {
	return nil
}
