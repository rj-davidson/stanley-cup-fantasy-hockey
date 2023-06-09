// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/entry"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/league"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/predicate"
)

// LeagueUpdate is the builder for updating League entities.
type LeagueUpdate struct {
	config
	hooks    []Hook
	mutation *LeagueMutation
}

// Where appends a list predicates to the LeagueUpdate builder.
func (lu *LeagueUpdate) Where(ps ...predicate.League) *LeagueUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetName sets the "name" field.
func (lu *LeagueUpdate) SetName(s string) *LeagueUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetPublic sets the "public" field.
func (lu *LeagueUpdate) SetPublic(b bool) *LeagueUpdate {
	lu.mutation.SetPublic(b)
	return lu
}

// SetNumForwards sets the "num_forwards" field.
func (lu *LeagueUpdate) SetNumForwards(i int) *LeagueUpdate {
	lu.mutation.ResetNumForwards()
	lu.mutation.SetNumForwards(i)
	return lu
}

// AddNumForwards adds i to the "num_forwards" field.
func (lu *LeagueUpdate) AddNumForwards(i int) *LeagueUpdate {
	lu.mutation.AddNumForwards(i)
	return lu
}

// SetNumDefenders sets the "num_defenders" field.
func (lu *LeagueUpdate) SetNumDefenders(i int) *LeagueUpdate {
	lu.mutation.ResetNumDefenders()
	lu.mutation.SetNumDefenders(i)
	return lu
}

// AddNumDefenders adds i to the "num_defenders" field.
func (lu *LeagueUpdate) AddNumDefenders(i int) *LeagueUpdate {
	lu.mutation.AddNumDefenders(i)
	return lu
}

// SetNumGoalies sets the "num_goalies" field.
func (lu *LeagueUpdate) SetNumGoalies(i int) *LeagueUpdate {
	lu.mutation.ResetNumGoalies()
	lu.mutation.SetNumGoalies(i)
	return lu
}

// AddNumGoalies adds i to the "num_goalies" field.
func (lu *LeagueUpdate) AddNumGoalies(i int) *LeagueUpdate {
	lu.mutation.AddNumGoalies(i)
	return lu
}

// AddEntryIDs adds the "entries" edge to the Entry entity by IDs.
func (lu *LeagueUpdate) AddEntryIDs(ids ...int) *LeagueUpdate {
	lu.mutation.AddEntryIDs(ids...)
	return lu
}

// AddEntries adds the "entries" edges to the Entry entity.
func (lu *LeagueUpdate) AddEntries(e ...*Entry) *LeagueUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return lu.AddEntryIDs(ids...)
}

// Mutation returns the LeagueMutation object of the builder.
func (lu *LeagueUpdate) Mutation() *LeagueMutation {
	return lu.mutation
}

// ClearEntries clears all "entries" edges to the Entry entity.
func (lu *LeagueUpdate) ClearEntries() *LeagueUpdate {
	lu.mutation.ClearEntries()
	return lu
}

// RemoveEntryIDs removes the "entries" edge to Entry entities by IDs.
func (lu *LeagueUpdate) RemoveEntryIDs(ids ...int) *LeagueUpdate {
	lu.mutation.RemoveEntryIDs(ids...)
	return lu
}

// RemoveEntries removes "entries" edges to Entry entities.
func (lu *LeagueUpdate) RemoveEntries(e ...*Entry) *LeagueUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return lu.RemoveEntryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LeagueUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LeagueUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LeagueUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LeagueUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LeagueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(league.Table, league.Columns, sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(league.FieldName, field.TypeString, value)
	}
	if value, ok := lu.mutation.Public(); ok {
		_spec.SetField(league.FieldPublic, field.TypeBool, value)
	}
	if value, ok := lu.mutation.NumForwards(); ok {
		_spec.SetField(league.FieldNumForwards, field.TypeInt, value)
	}
	if value, ok := lu.mutation.AddedNumForwards(); ok {
		_spec.AddField(league.FieldNumForwards, field.TypeInt, value)
	}
	if value, ok := lu.mutation.NumDefenders(); ok {
		_spec.SetField(league.FieldNumDefenders, field.TypeInt, value)
	}
	if value, ok := lu.mutation.AddedNumDefenders(); ok {
		_spec.AddField(league.FieldNumDefenders, field.TypeInt, value)
	}
	if value, ok := lu.mutation.NumGoalies(); ok {
		_spec.SetField(league.FieldNumGoalies, field.TypeInt, value)
	}
	if value, ok := lu.mutation.AddedNumGoalies(); ok {
		_spec.AddField(league.FieldNumGoalies, field.TypeInt, value)
	}
	if lu.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.EntriesTable,
			Columns: []string{league.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedEntriesIDs(); len(nodes) > 0 && !lu.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.EntriesTable,
			Columns: []string{league.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.EntriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.EntriesTable,
			Columns: []string{league.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{league.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LeagueUpdateOne is the builder for updating a single League entity.
type LeagueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LeagueMutation
}

// SetName sets the "name" field.
func (luo *LeagueUpdateOne) SetName(s string) *LeagueUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetPublic sets the "public" field.
func (luo *LeagueUpdateOne) SetPublic(b bool) *LeagueUpdateOne {
	luo.mutation.SetPublic(b)
	return luo
}

// SetNumForwards sets the "num_forwards" field.
func (luo *LeagueUpdateOne) SetNumForwards(i int) *LeagueUpdateOne {
	luo.mutation.ResetNumForwards()
	luo.mutation.SetNumForwards(i)
	return luo
}

// AddNumForwards adds i to the "num_forwards" field.
func (luo *LeagueUpdateOne) AddNumForwards(i int) *LeagueUpdateOne {
	luo.mutation.AddNumForwards(i)
	return luo
}

// SetNumDefenders sets the "num_defenders" field.
func (luo *LeagueUpdateOne) SetNumDefenders(i int) *LeagueUpdateOne {
	luo.mutation.ResetNumDefenders()
	luo.mutation.SetNumDefenders(i)
	return luo
}

// AddNumDefenders adds i to the "num_defenders" field.
func (luo *LeagueUpdateOne) AddNumDefenders(i int) *LeagueUpdateOne {
	luo.mutation.AddNumDefenders(i)
	return luo
}

// SetNumGoalies sets the "num_goalies" field.
func (luo *LeagueUpdateOne) SetNumGoalies(i int) *LeagueUpdateOne {
	luo.mutation.ResetNumGoalies()
	luo.mutation.SetNumGoalies(i)
	return luo
}

// AddNumGoalies adds i to the "num_goalies" field.
func (luo *LeagueUpdateOne) AddNumGoalies(i int) *LeagueUpdateOne {
	luo.mutation.AddNumGoalies(i)
	return luo
}

// AddEntryIDs adds the "entries" edge to the Entry entity by IDs.
func (luo *LeagueUpdateOne) AddEntryIDs(ids ...int) *LeagueUpdateOne {
	luo.mutation.AddEntryIDs(ids...)
	return luo
}

// AddEntries adds the "entries" edges to the Entry entity.
func (luo *LeagueUpdateOne) AddEntries(e ...*Entry) *LeagueUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return luo.AddEntryIDs(ids...)
}

// Mutation returns the LeagueMutation object of the builder.
func (luo *LeagueUpdateOne) Mutation() *LeagueMutation {
	return luo.mutation
}

// ClearEntries clears all "entries" edges to the Entry entity.
func (luo *LeagueUpdateOne) ClearEntries() *LeagueUpdateOne {
	luo.mutation.ClearEntries()
	return luo
}

// RemoveEntryIDs removes the "entries" edge to Entry entities by IDs.
func (luo *LeagueUpdateOne) RemoveEntryIDs(ids ...int) *LeagueUpdateOne {
	luo.mutation.RemoveEntryIDs(ids...)
	return luo
}

// RemoveEntries removes "entries" edges to Entry entities.
func (luo *LeagueUpdateOne) RemoveEntries(e ...*Entry) *LeagueUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return luo.RemoveEntryIDs(ids...)
}

// Where appends a list predicates to the LeagueUpdate builder.
func (luo *LeagueUpdateOne) Where(ps ...predicate.League) *LeagueUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LeagueUpdateOne) Select(field string, fields ...string) *LeagueUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated League entity.
func (luo *LeagueUpdateOne) Save(ctx context.Context) (*League, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LeagueUpdateOne) SaveX(ctx context.Context) *League {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LeagueUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LeagueUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LeagueUpdateOne) sqlSave(ctx context.Context) (_node *League, err error) {
	_spec := sqlgraph.NewUpdateSpec(league.Table, league.Columns, sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "League.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, league.FieldID)
		for _, f := range fields {
			if !league.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != league.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(league.FieldName, field.TypeString, value)
	}
	if value, ok := luo.mutation.Public(); ok {
		_spec.SetField(league.FieldPublic, field.TypeBool, value)
	}
	if value, ok := luo.mutation.NumForwards(); ok {
		_spec.SetField(league.FieldNumForwards, field.TypeInt, value)
	}
	if value, ok := luo.mutation.AddedNumForwards(); ok {
		_spec.AddField(league.FieldNumForwards, field.TypeInt, value)
	}
	if value, ok := luo.mutation.NumDefenders(); ok {
		_spec.SetField(league.FieldNumDefenders, field.TypeInt, value)
	}
	if value, ok := luo.mutation.AddedNumDefenders(); ok {
		_spec.AddField(league.FieldNumDefenders, field.TypeInt, value)
	}
	if value, ok := luo.mutation.NumGoalies(); ok {
		_spec.SetField(league.FieldNumGoalies, field.TypeInt, value)
	}
	if value, ok := luo.mutation.AddedNumGoalies(); ok {
		_spec.AddField(league.FieldNumGoalies, field.TypeInt, value)
	}
	if luo.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.EntriesTable,
			Columns: []string{league.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedEntriesIDs(); len(nodes) > 0 && !luo.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.EntriesTable,
			Columns: []string{league.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.EntriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.EntriesTable,
			Columns: []string{league.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &League{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{league.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
