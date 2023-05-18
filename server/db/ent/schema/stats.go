package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Stats holds the schema definition for the Stats entity.
type Stats struct {
	ent.Schema
}

// Fields of the Stats.
func (Stats) Fields() []ent.Field {
	return []ent.Field{
		field.Int("goals").Default(0),
		field.Int("assists").Default(0),
		field.Int("shutouts").Default(0),
		field.Int("wins").Default(0),
	}
}

// Edges of the Stats.
func (Stats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("player", Player.Type).
			Immutable().
			Unique(),
	}
}
