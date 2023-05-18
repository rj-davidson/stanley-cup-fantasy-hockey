// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/game"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/player"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/skaterstats"
)

// SkaterStatsCreate is the builder for creating a SkaterStats entity.
type SkaterStatsCreate struct {
	config
	mutation *SkaterStatsMutation
	hooks    []Hook
}

// SetGoals sets the "goals" field.
func (ssc *SkaterStatsCreate) SetGoals(i int) *SkaterStatsCreate {
	ssc.mutation.SetGoals(i)
	return ssc
}

// SetNillableGoals sets the "goals" field if the given value is not nil.
func (ssc *SkaterStatsCreate) SetNillableGoals(i *int) *SkaterStatsCreate {
	if i != nil {
		ssc.SetGoals(*i)
	}
	return ssc
}

// SetAssists sets the "assists" field.
func (ssc *SkaterStatsCreate) SetAssists(i int) *SkaterStatsCreate {
	ssc.mutation.SetAssists(i)
	return ssc
}

// SetNillableAssists sets the "assists" field if the given value is not nil.
func (ssc *SkaterStatsCreate) SetNillableAssists(i *int) *SkaterStatsCreate {
	if i != nil {
		ssc.SetAssists(*i)
	}
	return ssc
}

// SetHome sets the "home" field.
func (ssc *SkaterStatsCreate) SetHome(b bool) *SkaterStatsCreate {
	ssc.mutation.SetHome(b)
	return ssc
}

// SetNillableHome sets the "home" field if the given value is not nil.
func (ssc *SkaterStatsCreate) SetNillableHome(b *bool) *SkaterStatsCreate {
	if b != nil {
		ssc.SetHome(*b)
	}
	return ssc
}

// SetGameID sets the "game" edge to the Game entity by ID.
func (ssc *SkaterStatsCreate) SetGameID(id int) *SkaterStatsCreate {
	ssc.mutation.SetGameID(id)
	return ssc
}

// SetNillableGameID sets the "game" edge to the Game entity by ID if the given value is not nil.
func (ssc *SkaterStatsCreate) SetNillableGameID(id *int) *SkaterStatsCreate {
	if id != nil {
		ssc = ssc.SetGameID(*id)
	}
	return ssc
}

// SetGame sets the "game" edge to the Game entity.
func (ssc *SkaterStatsCreate) SetGame(g *Game) *SkaterStatsCreate {
	return ssc.SetGameID(g.ID)
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (ssc *SkaterStatsCreate) SetPlayerID(id int) *SkaterStatsCreate {
	ssc.mutation.SetPlayerID(id)
	return ssc
}

// SetNillablePlayerID sets the "player" edge to the Player entity by ID if the given value is not nil.
func (ssc *SkaterStatsCreate) SetNillablePlayerID(id *int) *SkaterStatsCreate {
	if id != nil {
		ssc = ssc.SetPlayerID(*id)
	}
	return ssc
}

// SetPlayer sets the "player" edge to the Player entity.
func (ssc *SkaterStatsCreate) SetPlayer(p *Player) *SkaterStatsCreate {
	return ssc.SetPlayerID(p.ID)
}

// Mutation returns the SkaterStatsMutation object of the builder.
func (ssc *SkaterStatsCreate) Mutation() *SkaterStatsMutation {
	return ssc.mutation
}

// Save creates the SkaterStats in the database.
func (ssc *SkaterStatsCreate) Save(ctx context.Context) (*SkaterStats, error) {
	ssc.defaults()
	return withHooks(ctx, ssc.sqlSave, ssc.mutation, ssc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ssc *SkaterStatsCreate) SaveX(ctx context.Context) *SkaterStats {
	v, err := ssc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ssc *SkaterStatsCreate) Exec(ctx context.Context) error {
	_, err := ssc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssc *SkaterStatsCreate) ExecX(ctx context.Context) {
	if err := ssc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ssc *SkaterStatsCreate) defaults() {
	if _, ok := ssc.mutation.Goals(); !ok {
		v := skaterstats.DefaultGoals
		ssc.mutation.SetGoals(v)
	}
	if _, ok := ssc.mutation.Assists(); !ok {
		v := skaterstats.DefaultAssists
		ssc.mutation.SetAssists(v)
	}
	if _, ok := ssc.mutation.Home(); !ok {
		v := skaterstats.DefaultHome
		ssc.mutation.SetHome(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ssc *SkaterStatsCreate) check() error {
	if _, ok := ssc.mutation.Goals(); !ok {
		return &ValidationError{Name: "goals", err: errors.New(`ent: missing required field "SkaterStats.goals"`)}
	}
	if _, ok := ssc.mutation.Assists(); !ok {
		return &ValidationError{Name: "assists", err: errors.New(`ent: missing required field "SkaterStats.assists"`)}
	}
	if _, ok := ssc.mutation.Home(); !ok {
		return &ValidationError{Name: "home", err: errors.New(`ent: missing required field "SkaterStats.home"`)}
	}
	return nil
}

func (ssc *SkaterStatsCreate) sqlSave(ctx context.Context) (*SkaterStats, error) {
	if err := ssc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ssc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ssc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ssc.mutation.id = &_node.ID
	ssc.mutation.done = true
	return _node, nil
}

func (ssc *SkaterStatsCreate) createSpec() (*SkaterStats, *sqlgraph.CreateSpec) {
	var (
		_node = &SkaterStats{config: ssc.config}
		_spec = sqlgraph.NewCreateSpec(skaterstats.Table, sqlgraph.NewFieldSpec(skaterstats.FieldID, field.TypeInt))
	)
	if value, ok := ssc.mutation.Goals(); ok {
		_spec.SetField(skaterstats.FieldGoals, field.TypeInt, value)
		_node.Goals = value
	}
	if value, ok := ssc.mutation.Assists(); ok {
		_spec.SetField(skaterstats.FieldAssists, field.TypeInt, value)
		_node.Assists = value
	}
	if value, ok := ssc.mutation.Home(); ok {
		_spec.SetField(skaterstats.FieldHome, field.TypeBool, value)
		_node.Home = value
	}
	if nodes := ssc.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   skaterstats.GameTable,
			Columns: []string{skaterstats.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.game_skater_stats = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ssc.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   skaterstats.PlayerTable,
			Columns: []string{skaterstats.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_skater_stats = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SkaterStatsCreateBulk is the builder for creating many SkaterStats entities in bulk.
type SkaterStatsCreateBulk struct {
	config
	builders []*SkaterStatsCreate
}

// Save creates the SkaterStats entities in the database.
func (sscb *SkaterStatsCreateBulk) Save(ctx context.Context) ([]*SkaterStats, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sscb.builders))
	nodes := make([]*SkaterStats, len(sscb.builders))
	mutators := make([]Mutator, len(sscb.builders))
	for i := range sscb.builders {
		func(i int, root context.Context) {
			builder := sscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SkaterStatsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sscb *SkaterStatsCreateBulk) SaveX(ctx context.Context) []*SkaterStats {
	v, err := sscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sscb *SkaterStatsCreateBulk) Exec(ctx context.Context) error {
	_, err := sscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sscb *SkaterStatsCreateBulk) ExecX(ctx context.Context) {
	if err := sscb.Exec(ctx); err != nil {
		panic(err)
	}
}
