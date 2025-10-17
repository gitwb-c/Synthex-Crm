package schema

import "entgo.io/ent"

// Costumer holds the schema definition for the Costumer entity.
type Costumer struct {
	ent.Schema
}

// Fields of the Costumer.
func (Costumer) Fields() []ent.Field {
	return nil
}

// Edges of the Costumer.
func (Costumer) Edges() []ent.Edge {
	return nil
}
