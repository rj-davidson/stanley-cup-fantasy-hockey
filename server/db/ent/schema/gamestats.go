package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GameStats holds the schema definition for the GameStats entity.
type GameStats struct {
	ent.Schema
}

// Fields of the GameStats.
func (GameStats) Fields() []ent.Field {
	return []ent.Field{
		field.Int("goals").Default(0),
		field.Int("assists").Default(0),
		field.Bool("win").Default(false),
		field.Bool("shutout").Default(false),
		field.Bool("homeGame").Default(false),
	}
}

// Edges of the GameStats.
func (GameStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game", Game.Type).Unique().Required().Immutable(),
		edge.To("player", Player.Type).Unique().Required().Immutable(),
	}
}
