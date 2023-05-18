// Code generated by ent, DO NOT EDIT.

package gamestats

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GameStats {
	return predicate.GameStats(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GameStats {
	return predicate.GameStats(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GameStats {
	return predicate.GameStats(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GameStats {
	return predicate.GameStats(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GameStats {
	return predicate.GameStats(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GameStats {
	return predicate.GameStats(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GameStats {
	return predicate.GameStats(sql.FieldLTE(FieldID, id))
}

// Goals applies equality check predicate on the "goals" field. It's identical to GoalsEQ.
func Goals(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldGoals, v))
}

// Assists applies equality check predicate on the "assists" field. It's identical to AssistsEQ.
func Assists(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldAssists, v))
}

// Win applies equality check predicate on the "win" field. It's identical to WinEQ.
func Win(v bool) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldWin, v))
}

// Shutout applies equality check predicate on the "shutout" field. It's identical to ShutoutEQ.
func Shutout(v bool) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldShutout, v))
}

// HomeGame applies equality check predicate on the "homeGame" field. It's identical to HomeGameEQ.
func HomeGame(v bool) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldHomeGame, v))
}

// GoalsEQ applies the EQ predicate on the "goals" field.
func GoalsEQ(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldGoals, v))
}

// GoalsNEQ applies the NEQ predicate on the "goals" field.
func GoalsNEQ(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldNEQ(FieldGoals, v))
}

// GoalsIn applies the In predicate on the "goals" field.
func GoalsIn(vs ...int) predicate.GameStats {
	return predicate.GameStats(sql.FieldIn(FieldGoals, vs...))
}

// GoalsNotIn applies the NotIn predicate on the "goals" field.
func GoalsNotIn(vs ...int) predicate.GameStats {
	return predicate.GameStats(sql.FieldNotIn(FieldGoals, vs...))
}

// GoalsGT applies the GT predicate on the "goals" field.
func GoalsGT(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldGT(FieldGoals, v))
}

// GoalsGTE applies the GTE predicate on the "goals" field.
func GoalsGTE(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldGTE(FieldGoals, v))
}

// GoalsLT applies the LT predicate on the "goals" field.
func GoalsLT(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldLT(FieldGoals, v))
}

// GoalsLTE applies the LTE predicate on the "goals" field.
func GoalsLTE(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldLTE(FieldGoals, v))
}

// AssistsEQ applies the EQ predicate on the "assists" field.
func AssistsEQ(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldAssists, v))
}

// AssistsNEQ applies the NEQ predicate on the "assists" field.
func AssistsNEQ(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldNEQ(FieldAssists, v))
}

// AssistsIn applies the In predicate on the "assists" field.
func AssistsIn(vs ...int) predicate.GameStats {
	return predicate.GameStats(sql.FieldIn(FieldAssists, vs...))
}

// AssistsNotIn applies the NotIn predicate on the "assists" field.
func AssistsNotIn(vs ...int) predicate.GameStats {
	return predicate.GameStats(sql.FieldNotIn(FieldAssists, vs...))
}

// AssistsGT applies the GT predicate on the "assists" field.
func AssistsGT(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldGT(FieldAssists, v))
}

// AssistsGTE applies the GTE predicate on the "assists" field.
func AssistsGTE(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldGTE(FieldAssists, v))
}

// AssistsLT applies the LT predicate on the "assists" field.
func AssistsLT(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldLT(FieldAssists, v))
}

// AssistsLTE applies the LTE predicate on the "assists" field.
func AssistsLTE(v int) predicate.GameStats {
	return predicate.GameStats(sql.FieldLTE(FieldAssists, v))
}

// WinEQ applies the EQ predicate on the "win" field.
func WinEQ(v bool) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldWin, v))
}

// WinNEQ applies the NEQ predicate on the "win" field.
func WinNEQ(v bool) predicate.GameStats {
	return predicate.GameStats(sql.FieldNEQ(FieldWin, v))
}

// ShutoutEQ applies the EQ predicate on the "shutout" field.
func ShutoutEQ(v bool) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldShutout, v))
}

// ShutoutNEQ applies the NEQ predicate on the "shutout" field.
func ShutoutNEQ(v bool) predicate.GameStats {
	return predicate.GameStats(sql.FieldNEQ(FieldShutout, v))
}

// HomeGameEQ applies the EQ predicate on the "homeGame" field.
func HomeGameEQ(v bool) predicate.GameStats {
	return predicate.GameStats(sql.FieldEQ(FieldHomeGame, v))
}

// HomeGameNEQ applies the NEQ predicate on the "homeGame" field.
func HomeGameNEQ(v bool) predicate.GameStats {
	return predicate.GameStats(sql.FieldNEQ(FieldHomeGame, v))
}

// HasGame applies the HasEdge predicate on the "game" edge.
func HasGame() predicate.GameStats {
	return predicate.GameStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, GameTable, GameColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGameWith applies the HasEdge predicate on the "game" edge with a given conditions (other predicates).
func HasGameWith(preds ...predicate.Game) predicate.GameStats {
	return predicate.GameStats(func(s *sql.Selector) {
		step := newGameStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlayer applies the HasEdge predicate on the "player" edge.
func HasPlayer() predicate.GameStats {
	return predicate.GameStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PlayerTable, PlayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlayerWith applies the HasEdge predicate on the "player" edge with a given conditions (other predicates).
func HasPlayerWith(preds ...predicate.Player) predicate.GameStats {
	return predicate.GameStats(func(s *sql.Selector) {
		step := newPlayerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GameStats) predicate.GameStats {
	return predicate.GameStats(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GameStats) predicate.GameStats {
	return predicate.GameStats(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GameStats) predicate.GameStats {
	return predicate.GameStats(func(s *sql.Selector) {
		p(s.Not())
	})
}