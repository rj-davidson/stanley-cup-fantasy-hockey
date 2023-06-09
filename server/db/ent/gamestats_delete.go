// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/gamestats"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/predicate"
)

// GameStatsDelete is the builder for deleting a GameStats entity.
type GameStatsDelete struct {
	config
	hooks    []Hook
	mutation *GameStatsMutation
}

// Where appends a list predicates to the GameStatsDelete builder.
func (gsd *GameStatsDelete) Where(ps ...predicate.GameStats) *GameStatsDelete {
	gsd.mutation.Where(ps...)
	return gsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gsd *GameStatsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gsd.sqlExec, gsd.mutation, gsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gsd *GameStatsDelete) ExecX(ctx context.Context) int {
	n, err := gsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gsd *GameStatsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(gamestats.Table, sqlgraph.NewFieldSpec(gamestats.FieldID, field.TypeInt))
	if ps := gsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gsd.mutation.done = true
	return affected, err
}

// GameStatsDeleteOne is the builder for deleting a single GameStats entity.
type GameStatsDeleteOne struct {
	gsd *GameStatsDelete
}

// Where appends a list predicates to the GameStatsDelete builder.
func (gsdo *GameStatsDeleteOne) Where(ps ...predicate.GameStats) *GameStatsDeleteOne {
	gsdo.gsd.mutation.Where(ps...)
	return gsdo
}

// Exec executes the deletion query.
func (gsdo *GameStatsDeleteOne) Exec(ctx context.Context) error {
	n, err := gsdo.gsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{gamestats.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gsdo *GameStatsDeleteOne) ExecX(ctx context.Context) {
	if err := gsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
