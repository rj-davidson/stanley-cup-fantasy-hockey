package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// League holds the schema definition for the League entity.
type League struct {
	ent.Schema
}

// Fields of the League.
func (League) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("season").Immutable(),
		field.Bool("public"),
		field.Int("num_forwards"),
		field.Int("num_defenders"),
		field.Int("num_goalies"),
		field.String("edit_key"),
		field.String("code").
			Unique(),
	}
}

// Edges of the League.
func (League) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("entries", Entry.Type),
	}
}
