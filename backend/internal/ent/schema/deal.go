package schema

import "entgo.io/ent"

// Deal holds the schema definition for the Deal entity.
type Deal struct {
	ent.Schema
}

// Fields of the Deal.
func (Deal) Fields() []ent.Field {
	return nil
}

// Edges of the Deal.
func (Deal) Edges() []ent.Edge {
	return nil
}
