package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Immutable(),
		field.String("name").Unique(),
		field.String("logo_filepath").Unique(),
		field.Bool("eliminated").Default(false),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("players", Player.Type),
		edge.To("homeGames", Game.Type),
		edge.To("awayGames", Game.Type),
	}
}
