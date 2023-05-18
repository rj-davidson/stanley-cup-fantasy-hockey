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
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/goaliestats"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/player"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/predicate"
)

// GoalieStatsUpdate is the builder for updating GoalieStats entities.
type GoalieStatsUpdate struct {
	config
	hooks    []Hook
	mutation *GoalieStatsMutation
}

// Where appends a list predicates to the GoalieStatsUpdate builder.
func (gsu *GoalieStatsUpdate) Where(ps ...predicate.GoalieStats) *GoalieStatsUpdate {
	gsu.mutation.Where(ps...)
	return gsu
}

// SetGoals sets the "goals" field.
func (gsu *GoalieStatsUpdate) SetGoals(i int) *GoalieStatsUpdate {
	gsu.mutation.ResetGoals()
	gsu.mutation.SetGoals(i)
	return gsu
}

// SetNillableGoals sets the "goals" field if the given value is not nil.
func (gsu *GoalieStatsUpdate) SetNillableGoals(i *int) *GoalieStatsUpdate {
	if i != nil {
		gsu.SetGoals(*i)
	}
	return gsu
}

// AddGoals adds i to the "goals" field.
func (gsu *GoalieStatsUpdate) AddGoals(i int) *GoalieStatsUpdate {
	gsu.mutation.AddGoals(i)
	return gsu
}

// SetAssists sets the "assists" field.
func (gsu *GoalieStatsUpdate) SetAssists(i int) *GoalieStatsUpdate {
	gsu.mutation.ResetAssists()
	gsu.mutation.SetAssists(i)
	return gsu
}

// SetNillableAssists sets the "assists" field if the given value is not nil.
func (gsu *GoalieStatsUpdate) SetNillableAssists(i *int) *GoalieStatsUpdate {
	if i != nil {
		gsu.SetAssists(*i)
	}
	return gsu
}

// AddAssists adds i to the "assists" field.
func (gsu *GoalieStatsUpdate) AddAssists(i int) *GoalieStatsUpdate {
	gsu.mutation.AddAssists(i)
	return gsu
}

// SetWin sets the "win" field.
func (gsu *GoalieStatsUpdate) SetWin(b bool) *GoalieStatsUpdate {
	gsu.mutation.SetWin(b)
	return gsu
}

// SetNillableWin sets the "win" field if the given value is not nil.
func (gsu *GoalieStatsUpdate) SetNillableWin(b *bool) *GoalieStatsUpdate {
	if b != nil {
		gsu.SetWin(*b)
	}
	return gsu
}

// SetLoss sets the "loss" field.
func (gsu *GoalieStatsUpdate) SetLoss(b bool) *GoalieStatsUpdate {
	gsu.mutation.SetLoss(b)
	return gsu
}

// SetNillableLoss sets the "loss" field if the given value is not nil.
func (gsu *GoalieStatsUpdate) SetNillableLoss(b *bool) *GoalieStatsUpdate {
	if b != nil {
		gsu.SetLoss(*b)
	}
	return gsu
}

// SetHome sets the "home" field.
func (gsu *GoalieStatsUpdate) SetHome(b bool) *GoalieStatsUpdate {
	gsu.mutation.SetHome(b)
	return gsu
}

// SetNillableHome sets the "home" field if the given value is not nil.
func (gsu *GoalieStatsUpdate) SetNillableHome(b *bool) *GoalieStatsUpdate {
	if b != nil {
		gsu.SetHome(*b)
	}
	return gsu
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (gsu *GoalieStatsUpdate) AddGameIDs(ids ...int) *GoalieStatsUpdate {
	gsu.mutation.AddGameIDs(ids...)
	return gsu
}

// AddGame adds the "game" edges to the Game entity.
func (gsu *GoalieStatsUpdate) AddGame(g ...*Game) *GoalieStatsUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gsu.AddGameIDs(ids...)
}

// AddPlayerIDs adds the "player" edge to the Player entity by IDs.
func (gsu *GoalieStatsUpdate) AddPlayerIDs(ids ...int) *GoalieStatsUpdate {
	gsu.mutation.AddPlayerIDs(ids...)
	return gsu
}

// AddPlayer adds the "player" edges to the Player entity.
func (gsu *GoalieStatsUpdate) AddPlayer(p ...*Player) *GoalieStatsUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gsu.AddPlayerIDs(ids...)
}

// Mutation returns the GoalieStatsMutation object of the builder.
func (gsu *GoalieStatsUpdate) Mutation() *GoalieStatsMutation {
	return gsu.mutation
}

// ClearGame clears all "game" edges to the Game entity.
func (gsu *GoalieStatsUpdate) ClearGame() *GoalieStatsUpdate {
	gsu.mutation.ClearGame()
	return gsu
}

// RemoveGameIDs removes the "game" edge to Game entities by IDs.
func (gsu *GoalieStatsUpdate) RemoveGameIDs(ids ...int) *GoalieStatsUpdate {
	gsu.mutation.RemoveGameIDs(ids...)
	return gsu
}

// RemoveGame removes "game" edges to Game entities.
func (gsu *GoalieStatsUpdate) RemoveGame(g ...*Game) *GoalieStatsUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gsu.RemoveGameIDs(ids...)
}

// ClearPlayer clears all "player" edges to the Player entity.
func (gsu *GoalieStatsUpdate) ClearPlayer() *GoalieStatsUpdate {
	gsu.mutation.ClearPlayer()
	return gsu
}

// RemovePlayerIDs removes the "player" edge to Player entities by IDs.
func (gsu *GoalieStatsUpdate) RemovePlayerIDs(ids ...int) *GoalieStatsUpdate {
	gsu.mutation.RemovePlayerIDs(ids...)
	return gsu
}

// RemovePlayer removes "player" edges to Player entities.
func (gsu *GoalieStatsUpdate) RemovePlayer(p ...*Player) *GoalieStatsUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gsu.RemovePlayerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gsu *GoalieStatsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gsu.sqlSave, gsu.mutation, gsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gsu *GoalieStatsUpdate) SaveX(ctx context.Context) int {
	affected, err := gsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gsu *GoalieStatsUpdate) Exec(ctx context.Context) error {
	_, err := gsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsu *GoalieStatsUpdate) ExecX(ctx context.Context) {
	if err := gsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gsu *GoalieStatsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(goaliestats.Table, goaliestats.Columns, sqlgraph.NewFieldSpec(goaliestats.FieldID, field.TypeInt))
	if ps := gsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gsu.mutation.Goals(); ok {
		_spec.SetField(goaliestats.FieldGoals, field.TypeInt, value)
	}
	if value, ok := gsu.mutation.AddedGoals(); ok {
		_spec.AddField(goaliestats.FieldGoals, field.TypeInt, value)
	}
	if value, ok := gsu.mutation.Assists(); ok {
		_spec.SetField(goaliestats.FieldAssists, field.TypeInt, value)
	}
	if value, ok := gsu.mutation.AddedAssists(); ok {
		_spec.AddField(goaliestats.FieldAssists, field.TypeInt, value)
	}
	if value, ok := gsu.mutation.Win(); ok {
		_spec.SetField(goaliestats.FieldWin, field.TypeBool, value)
	}
	if value, ok := gsu.mutation.Loss(); ok {
		_spec.SetField(goaliestats.FieldLoss, field.TypeBool, value)
	}
	if value, ok := gsu.mutation.Home(); ok {
		_spec.SetField(goaliestats.FieldHome, field.TypeBool, value)
	}
	if gsu.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.GameTable,
			Columns: goaliestats.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsu.mutation.RemovedGameIDs(); len(nodes) > 0 && !gsu.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.GameTable,
			Columns: goaliestats.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsu.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.GameTable,
			Columns: goaliestats.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gsu.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.PlayerTable,
			Columns: goaliestats.PlayerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsu.mutation.RemovedPlayerIDs(); len(nodes) > 0 && !gsu.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.PlayerTable,
			Columns: goaliestats.PlayerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsu.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.PlayerTable,
			Columns: goaliestats.PlayerPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, gsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goaliestats.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gsu.mutation.done = true
	return n, nil
}

// GoalieStatsUpdateOne is the builder for updating a single GoalieStats entity.
type GoalieStatsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoalieStatsMutation
}

// SetGoals sets the "goals" field.
func (gsuo *GoalieStatsUpdateOne) SetGoals(i int) *GoalieStatsUpdateOne {
	gsuo.mutation.ResetGoals()
	gsuo.mutation.SetGoals(i)
	return gsuo
}

// SetNillableGoals sets the "goals" field if the given value is not nil.
func (gsuo *GoalieStatsUpdateOne) SetNillableGoals(i *int) *GoalieStatsUpdateOne {
	if i != nil {
		gsuo.SetGoals(*i)
	}
	return gsuo
}

// AddGoals adds i to the "goals" field.
func (gsuo *GoalieStatsUpdateOne) AddGoals(i int) *GoalieStatsUpdateOne {
	gsuo.mutation.AddGoals(i)
	return gsuo
}

// SetAssists sets the "assists" field.
func (gsuo *GoalieStatsUpdateOne) SetAssists(i int) *GoalieStatsUpdateOne {
	gsuo.mutation.ResetAssists()
	gsuo.mutation.SetAssists(i)
	return gsuo
}

// SetNillableAssists sets the "assists" field if the given value is not nil.
func (gsuo *GoalieStatsUpdateOne) SetNillableAssists(i *int) *GoalieStatsUpdateOne {
	if i != nil {
		gsuo.SetAssists(*i)
	}
	return gsuo
}

// AddAssists adds i to the "assists" field.
func (gsuo *GoalieStatsUpdateOne) AddAssists(i int) *GoalieStatsUpdateOne {
	gsuo.mutation.AddAssists(i)
	return gsuo
}

// SetWin sets the "win" field.
func (gsuo *GoalieStatsUpdateOne) SetWin(b bool) *GoalieStatsUpdateOne {
	gsuo.mutation.SetWin(b)
	return gsuo
}

// SetNillableWin sets the "win" field if the given value is not nil.
func (gsuo *GoalieStatsUpdateOne) SetNillableWin(b *bool) *GoalieStatsUpdateOne {
	if b != nil {
		gsuo.SetWin(*b)
	}
	return gsuo
}

// SetLoss sets the "loss" field.
func (gsuo *GoalieStatsUpdateOne) SetLoss(b bool) *GoalieStatsUpdateOne {
	gsuo.mutation.SetLoss(b)
	return gsuo
}

// SetNillableLoss sets the "loss" field if the given value is not nil.
func (gsuo *GoalieStatsUpdateOne) SetNillableLoss(b *bool) *GoalieStatsUpdateOne {
	if b != nil {
		gsuo.SetLoss(*b)
	}
	return gsuo
}

// SetHome sets the "home" field.
func (gsuo *GoalieStatsUpdateOne) SetHome(b bool) *GoalieStatsUpdateOne {
	gsuo.mutation.SetHome(b)
	return gsuo
}

// SetNillableHome sets the "home" field if the given value is not nil.
func (gsuo *GoalieStatsUpdateOne) SetNillableHome(b *bool) *GoalieStatsUpdateOne {
	if b != nil {
		gsuo.SetHome(*b)
	}
	return gsuo
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (gsuo *GoalieStatsUpdateOne) AddGameIDs(ids ...int) *GoalieStatsUpdateOne {
	gsuo.mutation.AddGameIDs(ids...)
	return gsuo
}

// AddGame adds the "game" edges to the Game entity.
func (gsuo *GoalieStatsUpdateOne) AddGame(g ...*Game) *GoalieStatsUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gsuo.AddGameIDs(ids...)
}

// AddPlayerIDs adds the "player" edge to the Player entity by IDs.
func (gsuo *GoalieStatsUpdateOne) AddPlayerIDs(ids ...int) *GoalieStatsUpdateOne {
	gsuo.mutation.AddPlayerIDs(ids...)
	return gsuo
}

// AddPlayer adds the "player" edges to the Player entity.
func (gsuo *GoalieStatsUpdateOne) AddPlayer(p ...*Player) *GoalieStatsUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gsuo.AddPlayerIDs(ids...)
}

// Mutation returns the GoalieStatsMutation object of the builder.
func (gsuo *GoalieStatsUpdateOne) Mutation() *GoalieStatsMutation {
	return gsuo.mutation
}

// ClearGame clears all "game" edges to the Game entity.
func (gsuo *GoalieStatsUpdateOne) ClearGame() *GoalieStatsUpdateOne {
	gsuo.mutation.ClearGame()
	return gsuo
}

// RemoveGameIDs removes the "game" edge to Game entities by IDs.
func (gsuo *GoalieStatsUpdateOne) RemoveGameIDs(ids ...int) *GoalieStatsUpdateOne {
	gsuo.mutation.RemoveGameIDs(ids...)
	return gsuo
}

// RemoveGame removes "game" edges to Game entities.
func (gsuo *GoalieStatsUpdateOne) RemoveGame(g ...*Game) *GoalieStatsUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gsuo.RemoveGameIDs(ids...)
}

// ClearPlayer clears all "player" edges to the Player entity.
func (gsuo *GoalieStatsUpdateOne) ClearPlayer() *GoalieStatsUpdateOne {
	gsuo.mutation.ClearPlayer()
	return gsuo
}

// RemovePlayerIDs removes the "player" edge to Player entities by IDs.
func (gsuo *GoalieStatsUpdateOne) RemovePlayerIDs(ids ...int) *GoalieStatsUpdateOne {
	gsuo.mutation.RemovePlayerIDs(ids...)
	return gsuo
}

// RemovePlayer removes "player" edges to Player entities.
func (gsuo *GoalieStatsUpdateOne) RemovePlayer(p ...*Player) *GoalieStatsUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gsuo.RemovePlayerIDs(ids...)
}

// Where appends a list predicates to the GoalieStatsUpdate builder.
func (gsuo *GoalieStatsUpdateOne) Where(ps ...predicate.GoalieStats) *GoalieStatsUpdateOne {
	gsuo.mutation.Where(ps...)
	return gsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gsuo *GoalieStatsUpdateOne) Select(field string, fields ...string) *GoalieStatsUpdateOne {
	gsuo.fields = append([]string{field}, fields...)
	return gsuo
}

// Save executes the query and returns the updated GoalieStats entity.
func (gsuo *GoalieStatsUpdateOne) Save(ctx context.Context) (*GoalieStats, error) {
	return withHooks(ctx, gsuo.sqlSave, gsuo.mutation, gsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gsuo *GoalieStatsUpdateOne) SaveX(ctx context.Context) *GoalieStats {
	node, err := gsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gsuo *GoalieStatsUpdateOne) Exec(ctx context.Context) error {
	_, err := gsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsuo *GoalieStatsUpdateOne) ExecX(ctx context.Context) {
	if err := gsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gsuo *GoalieStatsUpdateOne) sqlSave(ctx context.Context) (_node *GoalieStats, err error) {
	_spec := sqlgraph.NewUpdateSpec(goaliestats.Table, goaliestats.Columns, sqlgraph.NewFieldSpec(goaliestats.FieldID, field.TypeInt))
	id, ok := gsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GoalieStats.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goaliestats.FieldID)
		for _, f := range fields {
			if !goaliestats.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goaliestats.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gsuo.mutation.Goals(); ok {
		_spec.SetField(goaliestats.FieldGoals, field.TypeInt, value)
	}
	if value, ok := gsuo.mutation.AddedGoals(); ok {
		_spec.AddField(goaliestats.FieldGoals, field.TypeInt, value)
	}
	if value, ok := gsuo.mutation.Assists(); ok {
		_spec.SetField(goaliestats.FieldAssists, field.TypeInt, value)
	}
	if value, ok := gsuo.mutation.AddedAssists(); ok {
		_spec.AddField(goaliestats.FieldAssists, field.TypeInt, value)
	}
	if value, ok := gsuo.mutation.Win(); ok {
		_spec.SetField(goaliestats.FieldWin, field.TypeBool, value)
	}
	if value, ok := gsuo.mutation.Loss(); ok {
		_spec.SetField(goaliestats.FieldLoss, field.TypeBool, value)
	}
	if value, ok := gsuo.mutation.Home(); ok {
		_spec.SetField(goaliestats.FieldHome, field.TypeBool, value)
	}
	if gsuo.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.GameTable,
			Columns: goaliestats.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsuo.mutation.RemovedGameIDs(); len(nodes) > 0 && !gsuo.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.GameTable,
			Columns: goaliestats.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsuo.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.GameTable,
			Columns: goaliestats.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gsuo.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.PlayerTable,
			Columns: goaliestats.PlayerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsuo.mutation.RemovedPlayerIDs(); len(nodes) > 0 && !gsuo.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.PlayerTable,
			Columns: goaliestats.PlayerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsuo.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goaliestats.PlayerTable,
			Columns: goaliestats.PlayerPrimaryKey,
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
	_node = &GoalieStats{config: gsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goaliestats.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gsuo.mutation.done = true
	return _node, nil
}
