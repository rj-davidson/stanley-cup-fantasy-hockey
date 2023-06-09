// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/team"
)

// Team is the model entity for the Team schema.
type Team struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LogoFilepath holds the value of the "logo_filepath" field.
	LogoFilepath string `json:"logo_filepath,omitempty"`
	// Eliminated holds the value of the "eliminated" field.
	Eliminated bool `json:"eliminated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeamQuery when eager-loading is set.
	Edges        TeamEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TeamEdges holds the relations/edges for other nodes in the graph.
type TeamEdges struct {
	// Players holds the value of the players edge.
	Players []*Player `json:"players,omitempty"`
	// HomeGames holds the value of the homeGames edge.
	HomeGames []*Game `json:"homeGames,omitempty"`
	// AwayGames holds the value of the awayGames edge.
	AwayGames []*Game `json:"awayGames,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PlayersOrErr returns the Players value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) PlayersOrErr() ([]*Player, error) {
	if e.loadedTypes[0] {
		return e.Players, nil
	}
	return nil, &NotLoadedError{edge: "players"}
}

// HomeGamesOrErr returns the HomeGames value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) HomeGamesOrErr() ([]*Game, error) {
	if e.loadedTypes[1] {
		return e.HomeGames, nil
	}
	return nil, &NotLoadedError{edge: "homeGames"}
}

// AwayGamesOrErr returns the AwayGames value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) AwayGamesOrErr() ([]*Game, error) {
	if e.loadedTypes[2] {
		return e.AwayGames, nil
	}
	return nil, &NotLoadedError{edge: "awayGames"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Team) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case team.FieldEliminated:
			values[i] = new(sql.NullBool)
		case team.FieldID:
			values[i] = new(sql.NullInt64)
		case team.FieldName, team.FieldLogoFilepath:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Team fields.
func (t *Team) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case team.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case team.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case team.FieldLogoFilepath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo_filepath", values[i])
			} else if value.Valid {
				t.LogoFilepath = value.String
			}
		case team.FieldEliminated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field eliminated", values[i])
			} else if value.Valid {
				t.Eliminated = value.Bool
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Team.
// This includes values selected through modifiers, order, etc.
func (t *Team) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryPlayers queries the "players" edge of the Team entity.
func (t *Team) QueryPlayers() *PlayerQuery {
	return NewTeamClient(t.config).QueryPlayers(t)
}

// QueryHomeGames queries the "homeGames" edge of the Team entity.
func (t *Team) QueryHomeGames() *GameQuery {
	return NewTeamClient(t.config).QueryHomeGames(t)
}

// QueryAwayGames queries the "awayGames" edge of the Team entity.
func (t *Team) QueryAwayGames() *GameQuery {
	return NewTeamClient(t.config).QueryAwayGames(t)
}

// Update returns a builder for updating this Team.
// Note that you need to call Team.Unwrap() before calling this method if this Team
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Team) Update() *TeamUpdateOne {
	return NewTeamClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Team entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Team) Unwrap() *Team {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Team is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Team) String() string {
	var builder strings.Builder
	builder.WriteString("Team(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("logo_filepath=")
	builder.WriteString(t.LogoFilepath)
	builder.WriteString(", ")
	builder.WriteString("eliminated=")
	builder.WriteString(fmt.Sprintf("%v", t.Eliminated))
	builder.WriteByte(')')
	return builder.String()
}

// Teams is a parsable slice of Team.
type Teams []*Team
