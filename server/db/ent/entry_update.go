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
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/player"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/predicate"
)

// EntryUpdate is the builder for updating Entry entities.
type EntryUpdate struct {
	config
	hooks    []Hook
	mutation *EntryMutation
}

// Where appends a list predicates to the EntryUpdate builder.
func (eu *EntryUpdate) Where(ps ...predicate.Entry) *EntryUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetOwnerName sets the "owner_name" field.
func (eu *EntryUpdate) SetOwnerName(s string) *EntryUpdate {
	eu.mutation.SetOwnerName(s)
	return eu
}

// SetLeagueID sets the "league" edge to the League entity by ID.
func (eu *EntryUpdate) SetLeagueID(id int) *EntryUpdate {
	eu.mutation.SetLeagueID(id)
	return eu
}

// SetNillableLeagueID sets the "league" edge to the League entity by ID if the given value is not nil.
func (eu *EntryUpdate) SetNillableLeagueID(id *int) *EntryUpdate {
	if id != nil {
		eu = eu.SetLeagueID(*id)
	}
	return eu
}

// SetLeague sets the "league" edge to the League entity.
func (eu *EntryUpdate) SetLeague(l *League) *EntryUpdate {
	return eu.SetLeagueID(l.ID)
}

// AddForwardIDs adds the "forwards" edge to the Player entity by IDs.
func (eu *EntryUpdate) AddForwardIDs(ids ...int) *EntryUpdate {
	eu.mutation.AddForwardIDs(ids...)
	return eu
}

// AddForwards adds the "forwards" edges to the Player entity.
func (eu *EntryUpdate) AddForwards(p ...*Player) *EntryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.AddForwardIDs(ids...)
}

// AddDefenderIDs adds the "defenders" edge to the Player entity by IDs.
func (eu *EntryUpdate) AddDefenderIDs(ids ...int) *EntryUpdate {
	eu.mutation.AddDefenderIDs(ids...)
	return eu
}

// AddDefenders adds the "defenders" edges to the Player entity.
func (eu *EntryUpdate) AddDefenders(p ...*Player) *EntryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.AddDefenderIDs(ids...)
}

// AddGoalyIDs adds the "goalies" edge to the Player entity by IDs.
func (eu *EntryUpdate) AddGoalyIDs(ids ...int) *EntryUpdate {
	eu.mutation.AddGoalyIDs(ids...)
	return eu
}

// AddGoalies adds the "goalies" edges to the Player entity.
func (eu *EntryUpdate) AddGoalies(p ...*Player) *EntryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.AddGoalyIDs(ids...)
}

// Mutation returns the EntryMutation object of the builder.
func (eu *EntryUpdate) Mutation() *EntryMutation {
	return eu.mutation
}

// ClearLeague clears the "league" edge to the League entity.
func (eu *EntryUpdate) ClearLeague() *EntryUpdate {
	eu.mutation.ClearLeague()
	return eu
}

// ClearForwards clears all "forwards" edges to the Player entity.
func (eu *EntryUpdate) ClearForwards() *EntryUpdate {
	eu.mutation.ClearForwards()
	return eu
}

// RemoveForwardIDs removes the "forwards" edge to Player entities by IDs.
func (eu *EntryUpdate) RemoveForwardIDs(ids ...int) *EntryUpdate {
	eu.mutation.RemoveForwardIDs(ids...)
	return eu
}

// RemoveForwards removes "forwards" edges to Player entities.
func (eu *EntryUpdate) RemoveForwards(p ...*Player) *EntryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.RemoveForwardIDs(ids...)
}

// ClearDefenders clears all "defenders" edges to the Player entity.
func (eu *EntryUpdate) ClearDefenders() *EntryUpdate {
	eu.mutation.ClearDefenders()
	return eu
}

// RemoveDefenderIDs removes the "defenders" edge to Player entities by IDs.
func (eu *EntryUpdate) RemoveDefenderIDs(ids ...int) *EntryUpdate {
	eu.mutation.RemoveDefenderIDs(ids...)
	return eu
}

// RemoveDefenders removes "defenders" edges to Player entities.
func (eu *EntryUpdate) RemoveDefenders(p ...*Player) *EntryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.RemoveDefenderIDs(ids...)
}

// ClearGoalies clears all "goalies" edges to the Player entity.
func (eu *EntryUpdate) ClearGoalies() *EntryUpdate {
	eu.mutation.ClearGoalies()
	return eu
}

// RemoveGoalyIDs removes the "goalies" edge to Player entities by IDs.
func (eu *EntryUpdate) RemoveGoalyIDs(ids ...int) *EntryUpdate {
	eu.mutation.RemoveGoalyIDs(ids...)
	return eu
}

// RemoveGoalies removes "goalies" edges to Player entities.
func (eu *EntryUpdate) RemoveGoalies(p ...*Player) *EntryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.RemoveGoalyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EntryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EntryUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EntryUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EntryUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EntryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(entry.Table, entry.Columns, sqlgraph.NewFieldSpec(entry.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.OwnerName(); ok {
		_spec.SetField(entry.FieldOwnerName, field.TypeString, value)
	}
	if eu.mutation.LeagueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entry.LeagueTable,
			Columns: []string{entry.LeagueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.LeagueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entry.LeagueTable,
			Columns: []string{entry.LeagueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.ForwardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.ForwardsTable,
			Columns: []string{entry.ForwardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedForwardsIDs(); len(nodes) > 0 && !eu.mutation.ForwardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.ForwardsTable,
			Columns: []string{entry.ForwardsColumn},
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
	if nodes := eu.mutation.ForwardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.ForwardsTable,
			Columns: []string{entry.ForwardsColumn},
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
	if eu.mutation.DefendersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.DefendersTable,
			Columns: []string{entry.DefendersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedDefendersIDs(); len(nodes) > 0 && !eu.mutation.DefendersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.DefendersTable,
			Columns: []string{entry.DefendersColumn},
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
	if nodes := eu.mutation.DefendersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.DefendersTable,
			Columns: []string{entry.DefendersColumn},
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
	if eu.mutation.GoaliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.GoaliesTable,
			Columns: []string{entry.GoaliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedGoaliesIDs(); len(nodes) > 0 && !eu.mutation.GoaliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.GoaliesTable,
			Columns: []string{entry.GoaliesColumn},
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
	if nodes := eu.mutation.GoaliesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.GoaliesTable,
			Columns: []string{entry.GoaliesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EntryUpdateOne is the builder for updating a single Entry entity.
type EntryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntryMutation
}

// SetOwnerName sets the "owner_name" field.
func (euo *EntryUpdateOne) SetOwnerName(s string) *EntryUpdateOne {
	euo.mutation.SetOwnerName(s)
	return euo
}

// SetLeagueID sets the "league" edge to the League entity by ID.
func (euo *EntryUpdateOne) SetLeagueID(id int) *EntryUpdateOne {
	euo.mutation.SetLeagueID(id)
	return euo
}

// SetNillableLeagueID sets the "league" edge to the League entity by ID if the given value is not nil.
func (euo *EntryUpdateOne) SetNillableLeagueID(id *int) *EntryUpdateOne {
	if id != nil {
		euo = euo.SetLeagueID(*id)
	}
	return euo
}

// SetLeague sets the "league" edge to the League entity.
func (euo *EntryUpdateOne) SetLeague(l *League) *EntryUpdateOne {
	return euo.SetLeagueID(l.ID)
}

// AddForwardIDs adds the "forwards" edge to the Player entity by IDs.
func (euo *EntryUpdateOne) AddForwardIDs(ids ...int) *EntryUpdateOne {
	euo.mutation.AddForwardIDs(ids...)
	return euo
}

// AddForwards adds the "forwards" edges to the Player entity.
func (euo *EntryUpdateOne) AddForwards(p ...*Player) *EntryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.AddForwardIDs(ids...)
}

// AddDefenderIDs adds the "defenders" edge to the Player entity by IDs.
func (euo *EntryUpdateOne) AddDefenderIDs(ids ...int) *EntryUpdateOne {
	euo.mutation.AddDefenderIDs(ids...)
	return euo
}

// AddDefenders adds the "defenders" edges to the Player entity.
func (euo *EntryUpdateOne) AddDefenders(p ...*Player) *EntryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.AddDefenderIDs(ids...)
}

// AddGoalyIDs adds the "goalies" edge to the Player entity by IDs.
func (euo *EntryUpdateOne) AddGoalyIDs(ids ...int) *EntryUpdateOne {
	euo.mutation.AddGoalyIDs(ids...)
	return euo
}

// AddGoalies adds the "goalies" edges to the Player entity.
func (euo *EntryUpdateOne) AddGoalies(p ...*Player) *EntryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.AddGoalyIDs(ids...)
}

// Mutation returns the EntryMutation object of the builder.
func (euo *EntryUpdateOne) Mutation() *EntryMutation {
	return euo.mutation
}

// ClearLeague clears the "league" edge to the League entity.
func (euo *EntryUpdateOne) ClearLeague() *EntryUpdateOne {
	euo.mutation.ClearLeague()
	return euo
}

// ClearForwards clears all "forwards" edges to the Player entity.
func (euo *EntryUpdateOne) ClearForwards() *EntryUpdateOne {
	euo.mutation.ClearForwards()
	return euo
}

// RemoveForwardIDs removes the "forwards" edge to Player entities by IDs.
func (euo *EntryUpdateOne) RemoveForwardIDs(ids ...int) *EntryUpdateOne {
	euo.mutation.RemoveForwardIDs(ids...)
	return euo
}

// RemoveForwards removes "forwards" edges to Player entities.
func (euo *EntryUpdateOne) RemoveForwards(p ...*Player) *EntryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.RemoveForwardIDs(ids...)
}

// ClearDefenders clears all "defenders" edges to the Player entity.
func (euo *EntryUpdateOne) ClearDefenders() *EntryUpdateOne {
	euo.mutation.ClearDefenders()
	return euo
}

// RemoveDefenderIDs removes the "defenders" edge to Player entities by IDs.
func (euo *EntryUpdateOne) RemoveDefenderIDs(ids ...int) *EntryUpdateOne {
	euo.mutation.RemoveDefenderIDs(ids...)
	return euo
}

// RemoveDefenders removes "defenders" edges to Player entities.
func (euo *EntryUpdateOne) RemoveDefenders(p ...*Player) *EntryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.RemoveDefenderIDs(ids...)
}

// ClearGoalies clears all "goalies" edges to the Player entity.
func (euo *EntryUpdateOne) ClearGoalies() *EntryUpdateOne {
	euo.mutation.ClearGoalies()
	return euo
}

// RemoveGoalyIDs removes the "goalies" edge to Player entities by IDs.
func (euo *EntryUpdateOne) RemoveGoalyIDs(ids ...int) *EntryUpdateOne {
	euo.mutation.RemoveGoalyIDs(ids...)
	return euo
}

// RemoveGoalies removes "goalies" edges to Player entities.
func (euo *EntryUpdateOne) RemoveGoalies(p ...*Player) *EntryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.RemoveGoalyIDs(ids...)
}

// Where appends a list predicates to the EntryUpdate builder.
func (euo *EntryUpdateOne) Where(ps ...predicate.Entry) *EntryUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EntryUpdateOne) Select(field string, fields ...string) *EntryUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Entry entity.
func (euo *EntryUpdateOne) Save(ctx context.Context) (*Entry, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EntryUpdateOne) SaveX(ctx context.Context) *Entry {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EntryUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EntryUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EntryUpdateOne) sqlSave(ctx context.Context) (_node *Entry, err error) {
	_spec := sqlgraph.NewUpdateSpec(entry.Table, entry.Columns, sqlgraph.NewFieldSpec(entry.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Entry.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entry.FieldID)
		for _, f := range fields {
			if !entry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.OwnerName(); ok {
		_spec.SetField(entry.FieldOwnerName, field.TypeString, value)
	}
	if euo.mutation.LeagueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entry.LeagueTable,
			Columns: []string{entry.LeagueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.LeagueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entry.LeagueTable,
			Columns: []string{entry.LeagueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.ForwardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.ForwardsTable,
			Columns: []string{entry.ForwardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedForwardsIDs(); len(nodes) > 0 && !euo.mutation.ForwardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.ForwardsTable,
			Columns: []string{entry.ForwardsColumn},
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
	if nodes := euo.mutation.ForwardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.ForwardsTable,
			Columns: []string{entry.ForwardsColumn},
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
	if euo.mutation.DefendersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.DefendersTable,
			Columns: []string{entry.DefendersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedDefendersIDs(); len(nodes) > 0 && !euo.mutation.DefendersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.DefendersTable,
			Columns: []string{entry.DefendersColumn},
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
	if nodes := euo.mutation.DefendersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.DefendersTable,
			Columns: []string{entry.DefendersColumn},
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
	if euo.mutation.GoaliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.GoaliesTable,
			Columns: []string{entry.GoaliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedGoaliesIDs(); len(nodes) > 0 && !euo.mutation.GoaliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.GoaliesTable,
			Columns: []string{entry.GoaliesColumn},
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
	if nodes := euo.mutation.GoaliesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entry.GoaliesTable,
			Columns: []string{entry.GoaliesColumn},
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
	_node = &Entry{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
