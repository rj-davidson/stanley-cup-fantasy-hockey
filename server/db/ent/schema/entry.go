package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Entry holds the schema definition for the Entry entity.
type Entry struct {
	ent.Schema
}

// Fields of the Entry.
func (Entry) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner_name"),
	}
}

// Edges of the Entry.
func (Entry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("league", League.Type).
			Unique().Ref("entries"),
		edge.To("players", Player.Type),
	}
}
