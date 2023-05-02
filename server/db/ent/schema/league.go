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
		field.Int("season").Immutable(),
		field.Bool("public"),
		field.Int("num_forwards"),
		field.Int("num_defenders"),
		field.Int("num_goalies"),
		field.Int("points_for_goal"),
		field.Int("points_for_assist"),
		field.Int("goalie_points_for_shutout"),
		field.Int("goalie_points_for_win"),
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
