package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Immutable(),
		field.String("name"),
		field.Enum("position").Values("Forward", "Defenseman", "Goalie"),
		field.Int("goals"),
		field.Int("assists"),
		field.Int("shutouts"),
		field.Int("wins"),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).Ref("players").Unique(),
		edge.From("entries", Entry.Type).Ref("players"),
		edge.To("homeGamesAsGoalie", Game.Type),
		edge.To("awayGamesAsGoalie", Game.Type),
	}
}
