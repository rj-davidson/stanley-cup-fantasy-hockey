// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/game"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/player"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/predicate"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/team"
)

// GameUpdate is the builder for updating Game entities.
type GameUpdate struct {
	config
	hooks    []Hook
	mutation *GameMutation
}

// Where appends a list predicates to the GameUpdate builder.
func (gu *GameUpdate) Where(ps ...predicate.Game) *GameUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetHomeScore sets the "homeScore" field.
func (gu *GameUpdate) SetHomeScore(i int) *GameUpdate {
	gu.mutation.ResetHomeScore()
	gu.mutation.SetHomeScore(i)
	return gu
}

// AddHomeScore adds i to the "homeScore" field.
func (gu *GameUpdate) AddHomeScore(i int) *GameUpdate {
	gu.mutation.AddHomeScore(i)
	return gu
}

// SetAwayScore sets the "awayScore" field.
func (gu *GameUpdate) SetAwayScore(i int) *GameUpdate {
	gu.mutation.ResetAwayScore()
	gu.mutation.SetAwayScore(i)
	return gu
}

// AddAwayScore adds i to the "awayScore" field.
func (gu *GameUpdate) AddAwayScore(i int) *GameUpdate {
	gu.mutation.AddAwayScore(i)
	return gu
}

// SetAwayTeamID sets the "awayTeam" edge to the Team entity by ID.
func (gu *GameUpdate) SetAwayTeamID(id int) *GameUpdate {
	gu.mutation.SetAwayTeamID(id)
	return gu
}

// SetAwayTeam sets the "awayTeam" edge to the Team entity.
func (gu *GameUpdate) SetAwayTeam(t *Team) *GameUpdate {
	return gu.SetAwayTeamID(t.ID)
}

// SetHomeTeamID sets the "homeTeam" edge to the Team entity by ID.
func (gu *GameUpdate) SetHomeTeamID(id int) *GameUpdate {
	gu.mutation.SetHomeTeamID(id)
	return gu
}

// SetHomeTeam sets the "homeTeam" edge to the Team entity.
func (gu *GameUpdate) SetHomeTeam(t *Team) *GameUpdate {
	return gu.SetHomeTeamID(t.ID)
}

// SetAwayGoalieID sets the "awayGoalie" edge to the Player entity by ID.
func (gu *GameUpdate) SetAwayGoalieID(id int) *GameUpdate {
	gu.mutation.SetAwayGoalieID(id)
	return gu
}

// SetAwayGoalie sets the "awayGoalie" edge to the Player entity.
func (gu *GameUpdate) SetAwayGoalie(p *Player) *GameUpdate {
	return gu.SetAwayGoalieID(p.ID)
}

// SetHomeGoalieID sets the "homeGoalie" edge to the Player entity by ID.
func (gu *GameUpdate) SetHomeGoalieID(id int) *GameUpdate {
	gu.mutation.SetHomeGoalieID(id)
	return gu
}

// SetHomeGoalie sets the "homeGoalie" edge to the Player entity.
func (gu *GameUpdate) SetHomeGoalie(p *Player) *GameUpdate {
	return gu.SetHomeGoalieID(p.ID)
}

// Mutation returns the GameMutation object of the builder.
func (gu *GameUpdate) Mutation() *GameMutation {
	return gu.mutation
}

// ClearAwayTeam clears the "awayTeam" edge to the Team entity.
func (gu *GameUpdate) ClearAwayTeam() *GameUpdate {
	gu.mutation.ClearAwayTeam()
	return gu
}

// ClearHomeTeam clears the "homeTeam" edge to the Team entity.
func (gu *GameUpdate) ClearHomeTeam() *GameUpdate {
	gu.mutation.ClearHomeTeam()
	return gu
}

// ClearAwayGoalie clears the "awayGoalie" edge to the Player entity.
func (gu *GameUpdate) ClearAwayGoalie() *GameUpdate {
	gu.mutation.ClearAwayGoalie()
	return gu
}

// ClearHomeGoalie clears the "homeGoalie" edge to the Player entity.
func (gu *GameUpdate) ClearHomeGoalie() *GameUpdate {
	gu.mutation.ClearHomeGoalie()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GameUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GameUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GameUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GameUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GameUpdate) check() error {
	if _, ok := gu.mutation.AwayTeamID(); gu.mutation.AwayTeamCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.awayTeam"`)
	}
	if _, ok := gu.mutation.HomeTeamID(); gu.mutation.HomeTeamCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.homeTeam"`)
	}
	if _, ok := gu.mutation.AwayGoalieID(); gu.mutation.AwayGoalieCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.awayGoalie"`)
	}
	if _, ok := gu.mutation.HomeGoalieID(); gu.mutation.HomeGoalieCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.homeGoalie"`)
	}
	return nil
}

func (gu *GameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(game.Table, game.Columns, sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.HomeScore(); ok {
		_spec.SetField(game.FieldHomeScore, field.TypeInt, value)
	}
	if value, ok := gu.mutation.AddedHomeScore(); ok {
		_spec.AddField(game.FieldHomeScore, field.TypeInt, value)
	}
	if value, ok := gu.mutation.AwayScore(); ok {
		_spec.SetField(game.FieldAwayScore, field.TypeInt, value)
	}
	if value, ok := gu.mutation.AddedAwayScore(); ok {
		_spec.AddField(game.FieldAwayScore, field.TypeInt, value)
	}
	if gu.mutation.AwayTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.AwayTeamTable,
			Columns: []string{game.AwayTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.AwayTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.AwayTeamTable,
			Columns: []string{game.AwayTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.HomeTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.HomeTeamTable,
			Columns: []string{game.HomeTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.HomeTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.HomeTeamTable,
			Columns: []string{game.HomeTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.AwayGoalieCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.AwayGoalieTable,
			Columns: []string{game.AwayGoalieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.AwayGoalieIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.AwayGoalieTable,
			Columns: []string{game.AwayGoalieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.HomeGoalieCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.HomeGoalieTable,
			Columns: []string{game.HomeGoalieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.HomeGoalieIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.HomeGoalieTable,
			Columns: []string{game.HomeGoalieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GameUpdateOne is the builder for updating a single Game entity.
type GameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameMutation
}

// SetHomeScore sets the "homeScore" field.
func (guo *GameUpdateOne) SetHomeScore(i int) *GameUpdateOne {
	guo.mutation.ResetHomeScore()
	guo.mutation.SetHomeScore(i)
	return guo
}

// AddHomeScore adds i to the "homeScore" field.
func (guo *GameUpdateOne) AddHomeScore(i int) *GameUpdateOne {
	guo.mutation.AddHomeScore(i)
	return guo
}

// SetAwayScore sets the "awayScore" field.
func (guo *GameUpdateOne) SetAwayScore(i int) *GameUpdateOne {
	guo.mutation.ResetAwayScore()
	guo.mutation.SetAwayScore(i)
	return guo
}

// AddAwayScore adds i to the "awayScore" field.
func (guo *GameUpdateOne) AddAwayScore(i int) *GameUpdateOne {
	guo.mutation.AddAwayScore(i)
	return guo
}

// SetAwayTeamID sets the "awayTeam" edge to the Team entity by ID.
func (guo *GameUpdateOne) SetAwayTeamID(id int) *GameUpdateOne {
	guo.mutation.SetAwayTeamID(id)
	return guo
}

// SetAwayTeam sets the "awayTeam" edge to the Team entity.
func (guo *GameUpdateOne) SetAwayTeam(t *Team) *GameUpdateOne {
	return guo.SetAwayTeamID(t.ID)
}

// SetHomeTeamID sets the "homeTeam" edge to the Team entity by ID.
func (guo *GameUpdateOne) SetHomeTeamID(id int) *GameUpdateOne {
	guo.mutation.SetHomeTeamID(id)
	return guo
}

// SetHomeTeam sets the "homeTeam" edge to the Team entity.
func (guo *GameUpdateOne) SetHomeTeam(t *Team) *GameUpdateOne {
	return guo.SetHomeTeamID(t.ID)
}

// SetAwayGoalieID sets the "awayGoalie" edge to the Player entity by ID.
func (guo *GameUpdateOne) SetAwayGoalieID(id int) *GameUpdateOne {
	guo.mutation.SetAwayGoalieID(id)
	return guo
}

// SetAwayGoalie sets the "awayGoalie" edge to the Player entity.
func (guo *GameUpdateOne) SetAwayGoalie(p *Player) *GameUpdateOne {
	return guo.SetAwayGoalieID(p.ID)
}

// SetHomeGoalieID sets the "homeGoalie" edge to the Player entity by ID.
func (guo *GameUpdateOne) SetHomeGoalieID(id int) *GameUpdateOne {
	guo.mutation.SetHomeGoalieID(id)
	return guo
}

// SetHomeGoalie sets the "homeGoalie" edge to the Player entity.
func (guo *GameUpdateOne) SetHomeGoalie(p *Player) *GameUpdateOne {
	return guo.SetHomeGoalieID(p.ID)
}

// Mutation returns the GameMutation object of the builder.
func (guo *GameUpdateOne) Mutation() *GameMutation {
	return guo.mutation
}

// ClearAwayTeam clears the "awayTeam" edge to the Team entity.
func (guo *GameUpdateOne) ClearAwayTeam() *GameUpdateOne {
	guo.mutation.ClearAwayTeam()
	return guo
}

// ClearHomeTeam clears the "homeTeam" edge to the Team entity.
func (guo *GameUpdateOne) ClearHomeTeam() *GameUpdateOne {
	guo.mutation.ClearHomeTeam()
	return guo
}

// ClearAwayGoalie clears the "awayGoalie" edge to the Player entity.
func (guo *GameUpdateOne) ClearAwayGoalie() *GameUpdateOne {
	guo.mutation.ClearAwayGoalie()
	return guo
}

// ClearHomeGoalie clears the "homeGoalie" edge to the Player entity.
func (guo *GameUpdateOne) ClearHomeGoalie() *GameUpdateOne {
	guo.mutation.ClearHomeGoalie()
	return guo
}

// Where appends a list predicates to the GameUpdate builder.
func (guo *GameUpdateOne) Where(ps ...predicate.Game) *GameUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GameUpdateOne) Select(field string, fields ...string) *GameUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Game entity.
func (guo *GameUpdateOne) Save(ctx context.Context) (*Game, error) {
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GameUpdateOne) SaveX(ctx context.Context) *Game {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GameUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GameUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GameUpdateOne) check() error {
	if _, ok := guo.mutation.AwayTeamID(); guo.mutation.AwayTeamCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.awayTeam"`)
	}
	if _, ok := guo.mutation.HomeTeamID(); guo.mutation.HomeTeamCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.homeTeam"`)
	}
	if _, ok := guo.mutation.AwayGoalieID(); guo.mutation.AwayGoalieCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.awayGoalie"`)
	}
	if _, ok := guo.mutation.HomeGoalieID(); guo.mutation.HomeGoalieCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.homeGoalie"`)
	}
	return nil
}

func (guo *GameUpdateOne) sqlSave(ctx context.Context) (_node *Game, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(game.Table, game.Columns, sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Game.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, game.FieldID)
		for _, f := range fields {
			if !game.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != game.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.HomeScore(); ok {
		_spec.SetField(game.FieldHomeScore, field.TypeInt, value)
	}
	if value, ok := guo.mutation.AddedHomeScore(); ok {
		_spec.AddField(game.FieldHomeScore, field.TypeInt, value)
	}
	if value, ok := guo.mutation.AwayScore(); ok {
		_spec.SetField(game.FieldAwayScore, field.TypeInt, value)
	}
	if value, ok := guo.mutation.AddedAwayScore(); ok {
		_spec.AddField(game.FieldAwayScore, field.TypeInt, value)
	}
	if guo.mutation.AwayTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.AwayTeamTable,
			Columns: []string{game.AwayTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.AwayTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.AwayTeamTable,
			Columns: []string{game.AwayTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.HomeTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.HomeTeamTable,
			Columns: []string{game.HomeTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.HomeTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.HomeTeamTable,
			Columns: []string{game.HomeTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.AwayGoalieCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.AwayGoalieTable,
			Columns: []string{game.AwayGoalieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.AwayGoalieIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.AwayGoalieTable,
			Columns: []string{game.AwayGoalieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.HomeGoalieCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.HomeGoalieTable,
			Columns: []string{game.HomeGoalieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.HomeGoalieIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.HomeGoalieTable,
			Columns: []string{game.HomeGoalieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Game{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
