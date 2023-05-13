package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Immutable(),
		field.Int("homeScore"),
		field.Int("awayScore"),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("awayTeam", Team.Type).
			Ref("awayGames").
			Unique().
			Required(),
		edge.From("homeTeam", Team.Type).
			Ref("homeGames").
			Unique().
			Required(),
		edge.From("awayGoalie", Player.Type).
			Ref("awayGamesAsGoalie").
			Unique().
			Required(),
		edge.From("homeGoalie", Player.Type).
			Ref("homeGamesAsGoalie").
			Unique().
			Required(),
	}
}
