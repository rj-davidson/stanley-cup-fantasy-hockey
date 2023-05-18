package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SkaterStats holds the schema definition for the SkaterStats entity.
type SkaterStats struct {
	ent.Schema
}

// Fields of the SkaterStats.
func (SkaterStats) Fields() []ent.Field {
	return []ent.Field{
		field.Int("goals").Default(0),
		field.Int("assists").Default(0),
		field.Bool("home").Default(false),
	}
}

// Edges of the SkaterStats.
func (SkaterStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("game", Game.Type).
			Ref("skaterStats").
			Unique(),
		edge.From("player", Player.Type).
			Ref("skaterStats").
			Unique(),
	}
}
