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
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/goaliestats"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/player"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/predicate"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/skaterstats"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/team"
)

// PlayerUpdate is the builder for updating Player entities.
type PlayerUpdate struct {
	config
	hooks    []Hook
	mutation *PlayerMutation
}

// Where appends a list predicates to the PlayerUpdate builder.
func (pu *PlayerUpdate) Where(ps ...predicate.Player) *PlayerUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PlayerUpdate) SetName(s string) *PlayerUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetPosition sets the "position" field.
func (pu *PlayerUpdate) SetPosition(pl player.Position) *PlayerUpdate {
	pu.mutation.SetPosition(pl)
	return pu
}

// SetTeamID sets the "team" edge to the Team entity by ID.
func (pu *PlayerUpdate) SetTeamID(id int) *PlayerUpdate {
	pu.mutation.SetTeamID(id)
	return pu
}

// SetNillableTeamID sets the "team" edge to the Team entity by ID if the given value is not nil.
func (pu *PlayerUpdate) SetNillableTeamID(id *int) *PlayerUpdate {
	if id != nil {
		pu = pu.SetTeamID(*id)
	}
	return pu
}

// SetTeam sets the "team" edge to the Team entity.
func (pu *PlayerUpdate) SetTeam(t *Team) *PlayerUpdate {
	return pu.SetTeamID(t.ID)
}

// AddEntryIDs adds the "entries" edge to the Entry entity by IDs.
func (pu *PlayerUpdate) AddEntryIDs(ids ...int) *PlayerUpdate {
	pu.mutation.AddEntryIDs(ids...)
	return pu
}

// AddEntries adds the "entries" edges to the Entry entity.
func (pu *PlayerUpdate) AddEntries(e ...*Entry) *PlayerUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pu.AddEntryIDs(ids...)
}

// AddSkaterStatIDs adds the "skaterStats" edge to the SkaterStats entity by IDs.
func (pu *PlayerUpdate) AddSkaterStatIDs(ids ...int) *PlayerUpdate {
	pu.mutation.AddSkaterStatIDs(ids...)
	return pu
}

// AddSkaterStats adds the "skaterStats" edges to the SkaterStats entity.
func (pu *PlayerUpdate) AddSkaterStats(s ...*SkaterStats) *PlayerUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddSkaterStatIDs(ids...)
}

// AddGoalieStatIDs adds the "goalieStats" edge to the GoalieStats entity by IDs.
func (pu *PlayerUpdate) AddGoalieStatIDs(ids ...int) *PlayerUpdate {
	pu.mutation.AddGoalieStatIDs(ids...)
	return pu
}

// AddGoalieStats adds the "goalieStats" edges to the GoalieStats entity.
func (pu *PlayerUpdate) AddGoalieStats(g ...*GoalieStats) *PlayerUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return pu.AddGoalieStatIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (pu *PlayerUpdate) Mutation() *PlayerMutation {
	return pu.mutation
}

// ClearTeam clears the "team" edge to the Team entity.
func (pu *PlayerUpdate) ClearTeam() *PlayerUpdate {
	pu.mutation.ClearTeam()
	return pu
}

// ClearEntries clears all "entries" edges to the Entry entity.
func (pu *PlayerUpdate) ClearEntries() *PlayerUpdate {
	pu.mutation.ClearEntries()
	return pu
}

// RemoveEntryIDs removes the "entries" edge to Entry entities by IDs.
func (pu *PlayerUpdate) RemoveEntryIDs(ids ...int) *PlayerUpdate {
	pu.mutation.RemoveEntryIDs(ids...)
	return pu
}

// RemoveEntries removes "entries" edges to Entry entities.
func (pu *PlayerUpdate) RemoveEntries(e ...*Entry) *PlayerUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pu.RemoveEntryIDs(ids...)
}

// ClearSkaterStats clears all "skaterStats" edges to the SkaterStats entity.
func (pu *PlayerUpdate) ClearSkaterStats() *PlayerUpdate {
	pu.mutation.ClearSkaterStats()
	return pu
}

// RemoveSkaterStatIDs removes the "skaterStats" edge to SkaterStats entities by IDs.
func (pu *PlayerUpdate) RemoveSkaterStatIDs(ids ...int) *PlayerUpdate {
	pu.mutation.RemoveSkaterStatIDs(ids...)
	return pu
}

// RemoveSkaterStats removes "skaterStats" edges to SkaterStats entities.
func (pu *PlayerUpdate) RemoveSkaterStats(s ...*SkaterStats) *PlayerUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveSkaterStatIDs(ids...)
}

// ClearGoalieStats clears all "goalieStats" edges to the GoalieStats entity.
func (pu *PlayerUpdate) ClearGoalieStats() *PlayerUpdate {
	pu.mutation.ClearGoalieStats()
	return pu
}

// RemoveGoalieStatIDs removes the "goalieStats" edge to GoalieStats entities by IDs.
func (pu *PlayerUpdate) RemoveGoalieStatIDs(ids ...int) *PlayerUpdate {
	pu.mutation.RemoveGoalieStatIDs(ids...)
	return pu
}

// RemoveGoalieStats removes "goalieStats" edges to GoalieStats entities.
func (pu *PlayerUpdate) RemoveGoalieStats(g ...*GoalieStats) *PlayerUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return pu.RemoveGoalieStatIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlayerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlayerUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlayerUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlayerUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PlayerUpdate) check() error {
	if v, ok := pu.mutation.Position(); ok {
		if err := player.PositionValidator(v); err != nil {
			return &ValidationError{Name: "position", err: fmt.Errorf(`ent: validator failed for field "Player.position": %w`, err)}
		}
	}
	return nil
}

func (pu *PlayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(player.Table, player.Columns, sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(player.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Position(); ok {
		_spec.SetField(player.FieldPosition, field.TypeEnum, value)
	}
	if pu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   player.TeamTable,
			Columns: []string{player.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   player.TeamTable,
			Columns: []string{player.TeamColumn},
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
	if pu.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.EntriesTable,
			Columns: player.EntriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedEntriesIDs(); len(nodes) > 0 && !pu.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.EntriesTable,
			Columns: player.EntriesPrimaryKey,
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
	if nodes := pu.mutation.EntriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.EntriesTable,
			Columns: player.EntriesPrimaryKey,
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
	if pu.mutation.SkaterStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.SkaterStatsTable,
			Columns: []string{player.SkaterStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skaterstats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedSkaterStatsIDs(); len(nodes) > 0 && !pu.mutation.SkaterStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.SkaterStatsTable,
			Columns: []string{player.SkaterStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skaterstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SkaterStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.SkaterStatsTable,
			Columns: []string{player.SkaterStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skaterstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.GoalieStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.GoalieStatsTable,
			Columns: player.GoalieStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(goaliestats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedGoalieStatsIDs(); len(nodes) > 0 && !pu.mutation.GoalieStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.GoalieStatsTable,
			Columns: player.GoalieStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(goaliestats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.GoalieStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.GoalieStatsTable,
			Columns: player.GoalieStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(goaliestats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PlayerUpdateOne is the builder for updating a single Player entity.
type PlayerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlayerMutation
}

// SetName sets the "name" field.
func (puo *PlayerUpdateOne) SetName(s string) *PlayerUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetPosition sets the "position" field.
func (puo *PlayerUpdateOne) SetPosition(pl player.Position) *PlayerUpdateOne {
	puo.mutation.SetPosition(pl)
	return puo
}

// SetTeamID sets the "team" edge to the Team entity by ID.
func (puo *PlayerUpdateOne) SetTeamID(id int) *PlayerUpdateOne {
	puo.mutation.SetTeamID(id)
	return puo
}

// SetNillableTeamID sets the "team" edge to the Team entity by ID if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableTeamID(id *int) *PlayerUpdateOne {
	if id != nil {
		puo = puo.SetTeamID(*id)
	}
	return puo
}

// SetTeam sets the "team" edge to the Team entity.
func (puo *PlayerUpdateOne) SetTeam(t *Team) *PlayerUpdateOne {
	return puo.SetTeamID(t.ID)
}

// AddEntryIDs adds the "entries" edge to the Entry entity by IDs.
func (puo *PlayerUpdateOne) AddEntryIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.AddEntryIDs(ids...)
	return puo
}

// AddEntries adds the "entries" edges to the Entry entity.
func (puo *PlayerUpdateOne) AddEntries(e ...*Entry) *PlayerUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return puo.AddEntryIDs(ids...)
}

// AddSkaterStatIDs adds the "skaterStats" edge to the SkaterStats entity by IDs.
func (puo *PlayerUpdateOne) AddSkaterStatIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.AddSkaterStatIDs(ids...)
	return puo
}

// AddSkaterStats adds the "skaterStats" edges to the SkaterStats entity.
func (puo *PlayerUpdateOne) AddSkaterStats(s ...*SkaterStats) *PlayerUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddSkaterStatIDs(ids...)
}

// AddGoalieStatIDs adds the "goalieStats" edge to the GoalieStats entity by IDs.
func (puo *PlayerUpdateOne) AddGoalieStatIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.AddGoalieStatIDs(ids...)
	return puo
}

// AddGoalieStats adds the "goalieStats" edges to the GoalieStats entity.
func (puo *PlayerUpdateOne) AddGoalieStats(g ...*GoalieStats) *PlayerUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return puo.AddGoalieStatIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (puo *PlayerUpdateOne) Mutation() *PlayerMutation {
	return puo.mutation
}

// ClearTeam clears the "team" edge to the Team entity.
func (puo *PlayerUpdateOne) ClearTeam() *PlayerUpdateOne {
	puo.mutation.ClearTeam()
	return puo
}

// ClearEntries clears all "entries" edges to the Entry entity.
func (puo *PlayerUpdateOne) ClearEntries() *PlayerUpdateOne {
	puo.mutation.ClearEntries()
	return puo
}

// RemoveEntryIDs removes the "entries" edge to Entry entities by IDs.
func (puo *PlayerUpdateOne) RemoveEntryIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.RemoveEntryIDs(ids...)
	return puo
}

// RemoveEntries removes "entries" edges to Entry entities.
func (puo *PlayerUpdateOne) RemoveEntries(e ...*Entry) *PlayerUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return puo.RemoveEntryIDs(ids...)
}

// ClearSkaterStats clears all "skaterStats" edges to the SkaterStats entity.
func (puo *PlayerUpdateOne) ClearSkaterStats() *PlayerUpdateOne {
	puo.mutation.ClearSkaterStats()
	return puo
}

// RemoveSkaterStatIDs removes the "skaterStats" edge to SkaterStats entities by IDs.
func (puo *PlayerUpdateOne) RemoveSkaterStatIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.RemoveSkaterStatIDs(ids...)
	return puo
}

// RemoveSkaterStats removes "skaterStats" edges to SkaterStats entities.
func (puo *PlayerUpdateOne) RemoveSkaterStats(s ...*SkaterStats) *PlayerUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveSkaterStatIDs(ids...)
}

// ClearGoalieStats clears all "goalieStats" edges to the GoalieStats entity.
func (puo *PlayerUpdateOne) ClearGoalieStats() *PlayerUpdateOne {
	puo.mutation.ClearGoalieStats()
	return puo
}

// RemoveGoalieStatIDs removes the "goalieStats" edge to GoalieStats entities by IDs.
func (puo *PlayerUpdateOne) RemoveGoalieStatIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.RemoveGoalieStatIDs(ids...)
	return puo
}

// RemoveGoalieStats removes "goalieStats" edges to GoalieStats entities.
func (puo *PlayerUpdateOne) RemoveGoalieStats(g ...*GoalieStats) *PlayerUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return puo.RemoveGoalieStatIDs(ids...)
}

// Where appends a list predicates to the PlayerUpdate builder.
func (puo *PlayerUpdateOne) Where(ps ...predicate.Player) *PlayerUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlayerUpdateOne) Select(field string, fields ...string) *PlayerUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Player entity.
func (puo *PlayerUpdateOne) Save(ctx context.Context) (*Player, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlayerUpdateOne) SaveX(ctx context.Context) *Player {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlayerUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlayerUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PlayerUpdateOne) check() error {
	if v, ok := puo.mutation.Position(); ok {
		if err := player.PositionValidator(v); err != nil {
			return &ValidationError{Name: "position", err: fmt.Errorf(`ent: validator failed for field "Player.position": %w`, err)}
		}
	}
	return nil
}

func (puo *PlayerUpdateOne) sqlSave(ctx context.Context) (_node *Player, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(player.Table, player.Columns, sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Player.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, player.FieldID)
		for _, f := range fields {
			if !player.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != player.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(player.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Position(); ok {
		_spec.SetField(player.FieldPosition, field.TypeEnum, value)
	}
	if puo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   player.TeamTable,
			Columns: []string{player.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   player.TeamTable,
			Columns: []string{player.TeamColumn},
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
	if puo.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.EntriesTable,
			Columns: player.EntriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedEntriesIDs(); len(nodes) > 0 && !puo.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.EntriesTable,
			Columns: player.EntriesPrimaryKey,
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
	if nodes := puo.mutation.EntriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.EntriesTable,
			Columns: player.EntriesPrimaryKey,
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
	if puo.mutation.SkaterStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.SkaterStatsTable,
			Columns: []string{player.SkaterStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skaterstats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedSkaterStatsIDs(); len(nodes) > 0 && !puo.mutation.SkaterStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.SkaterStatsTable,
			Columns: []string{player.SkaterStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skaterstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SkaterStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.SkaterStatsTable,
			Columns: []string{player.SkaterStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skaterstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.GoalieStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.GoalieStatsTable,
			Columns: player.GoalieStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(goaliestats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedGoalieStatsIDs(); len(nodes) > 0 && !puo.mutation.GoalieStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.GoalieStatsTable,
			Columns: player.GoalieStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(goaliestats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.GoalieStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   player.GoalieStatsTable,
			Columns: player.GoalieStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(goaliestats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Player{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
