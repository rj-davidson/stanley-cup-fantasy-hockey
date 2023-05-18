package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GoalieStats holds the schema definition for the GoalieStats entity.
type GoalieStats struct {
	ent.Schema
}

// Fields of the GoalieStats.
func (GoalieStats) Fields() []ent.Field {
	return []ent.Field{
		field.Int("goals").Default(0),
		field.Int("assists").Default(0),
		field.Bool("win").Default(false),
		field.Bool("loss").Default(false),
		field.Bool("home").Default(false),
	}
}

// Edges of the GoalieStats.
func (GoalieStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("game", Game.Type).
			Ref("goalieStats"),
		edge.From("player", Player.Type).
			Ref("goalieStats"),
	}
}
