// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/game"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/team"
)

// Game is the model entity for the Game schema.
type Game struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// HomeScore holds the value of the "homeScore" field.
	HomeScore int `json:"homeScore,omitempty"`
	// AwayScore holds the value of the "awayScore" field.
	AwayScore int `json:"awayScore,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GameQuery when eager-loading is set.
	Edges           GameEdges `json:"edges"`
	team_home_games *int
	team_away_games *int
	selectValues    sql.SelectValues
}

// GameEdges holds the relations/edges for other nodes in the graph.
type GameEdges struct {
	// AwayTeam holds the value of the awayTeam edge.
	AwayTeam *Team `json:"awayTeam,omitempty"`
	// HomeTeam holds the value of the homeTeam edge.
	HomeTeam *Team `json:"homeTeam,omitempty"`
	// SkaterStats holds the value of the skaterStats edge.
	SkaterStats []*SkaterStats `json:"skaterStats,omitempty"`
	// GoalieStats holds the value of the goalieStats edge.
	GoalieStats []*GoalieStats `json:"goalieStats,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// AwayTeamOrErr returns the AwayTeam value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameEdges) AwayTeamOrErr() (*Team, error) {
	if e.loadedTypes[0] {
		if e.AwayTeam == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.AwayTeam, nil
	}
	return nil, &NotLoadedError{edge: "awayTeam"}
}

// HomeTeamOrErr returns the HomeTeam value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameEdges) HomeTeamOrErr() (*Team, error) {
	if e.loadedTypes[1] {
		if e.HomeTeam == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.HomeTeam, nil
	}
	return nil, &NotLoadedError{edge: "homeTeam"}
}

// SkaterStatsOrErr returns the SkaterStats value or an error if the edge
// was not loaded in eager-loading.
func (e GameEdges) SkaterStatsOrErr() ([]*SkaterStats, error) {
	if e.loadedTypes[2] {
		return e.SkaterStats, nil
	}
	return nil, &NotLoadedError{edge: "skaterStats"}
}

// GoalieStatsOrErr returns the GoalieStats value or an error if the edge
// was not loaded in eager-loading.
func (e GameEdges) GoalieStatsOrErr() ([]*GoalieStats, error) {
	if e.loadedTypes[3] {
		return e.GoalieStats, nil
	}
	return nil, &NotLoadedError{edge: "goalieStats"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Game) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case game.FieldID, game.FieldHomeScore, game.FieldAwayScore:
			values[i] = new(sql.NullInt64)
		case game.ForeignKeys[0]: // team_home_games
			values[i] = new(sql.NullInt64)
		case game.ForeignKeys[1]: // team_away_games
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Game fields.
func (ga *Game) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case game.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ga.ID = int(value.Int64)
		case game.FieldHomeScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field homeScore", values[i])
			} else if value.Valid {
				ga.HomeScore = int(value.Int64)
			}
		case game.FieldAwayScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field awayScore", values[i])
			} else if value.Valid {
				ga.AwayScore = int(value.Int64)
			}
		case game.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field team_home_games", value)
			} else if value.Valid {
				ga.team_home_games = new(int)
				*ga.team_home_games = int(value.Int64)
			}
		case game.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field team_away_games", value)
			} else if value.Valid {
				ga.team_away_games = new(int)
				*ga.team_away_games = int(value.Int64)
			}
		default:
			ga.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Game.
// This includes values selected through modifiers, order, etc.
func (ga *Game) Value(name string) (ent.Value, error) {
	return ga.selectValues.Get(name)
}

// QueryAwayTeam queries the "awayTeam" edge of the Game entity.
func (ga *Game) QueryAwayTeam() *TeamQuery {
	return NewGameClient(ga.config).QueryAwayTeam(ga)
}

// QueryHomeTeam queries the "homeTeam" edge of the Game entity.
func (ga *Game) QueryHomeTeam() *TeamQuery {
	return NewGameClient(ga.config).QueryHomeTeam(ga)
}

// QuerySkaterStats queries the "skaterStats" edge of the Game entity.
func (ga *Game) QuerySkaterStats() *SkaterStatsQuery {
	return NewGameClient(ga.config).QuerySkaterStats(ga)
}

// QueryGoalieStats queries the "goalieStats" edge of the Game entity.
func (ga *Game) QueryGoalieStats() *GoalieStatsQuery {
	return NewGameClient(ga.config).QueryGoalieStats(ga)
}

// Update returns a builder for updating this Game.
// Note that you need to call Game.Unwrap() before calling this method if this Game
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *Game) Update() *GameUpdateOne {
	return NewGameClient(ga.config).UpdateOne(ga)
}

// Unwrap unwraps the Game entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *Game) Unwrap() *Game {
	_tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: Game is not a transactional entity")
	}
	ga.config.driver = _tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *Game) String() string {
	var builder strings.Builder
	builder.WriteString("Game(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ga.ID))
	builder.WriteString("homeScore=")
	builder.WriteString(fmt.Sprintf("%v", ga.HomeScore))
	builder.WriteString(", ")
	builder.WriteString("awayScore=")
	builder.WriteString(fmt.Sprintf("%v", ga.AwayScore))
	builder.WriteByte(')')
	return builder.String()
}

// Games is a parsable slice of Game.
type Games []*Game
