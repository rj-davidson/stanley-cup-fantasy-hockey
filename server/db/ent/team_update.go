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

// TeamUpdate is the builder for updating Team entities.
type TeamUpdate struct {
	config
	hooks    []Hook
	mutation *TeamMutation
}

// Where appends a list predicates to the TeamUpdate builder.
func (tu *TeamUpdate) Where(ps ...predicate.Team) *TeamUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TeamUpdate) SetName(s string) *TeamUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetLogoFilepath sets the "logo_filepath" field.
func (tu *TeamUpdate) SetLogoFilepath(s string) *TeamUpdate {
	tu.mutation.SetLogoFilepath(s)
	return tu
}

// SetEliminated sets the "eliminated" field.
func (tu *TeamUpdate) SetEliminated(b bool) *TeamUpdate {
	tu.mutation.SetEliminated(b)
	return tu
}

// SetNillableEliminated sets the "eliminated" field if the given value is not nil.
func (tu *TeamUpdate) SetNillableEliminated(b *bool) *TeamUpdate {
	if b != nil {
		tu.SetEliminated(*b)
	}
	return tu
}

// AddPlayerIDs adds the "players" edge to the Player entity by IDs.
func (tu *TeamUpdate) AddPlayerIDs(ids ...int) *TeamUpdate {
	tu.mutation.AddPlayerIDs(ids...)
	return tu
}

// AddPlayers adds the "players" edges to the Player entity.
func (tu *TeamUpdate) AddPlayers(p ...*Player) *TeamUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.AddPlayerIDs(ids...)
}

// AddHomeGameIDs adds the "homeGames" edge to the Game entity by IDs.
func (tu *TeamUpdate) AddHomeGameIDs(ids ...int) *TeamUpdate {
	tu.mutation.AddHomeGameIDs(ids...)
	return tu
}

// AddHomeGames adds the "homeGames" edges to the Game entity.
func (tu *TeamUpdate) AddHomeGames(g ...*Game) *TeamUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.AddHomeGameIDs(ids...)
}

// AddAwayGameIDs adds the "awayGames" edge to the Game entity by IDs.
func (tu *TeamUpdate) AddAwayGameIDs(ids ...int) *TeamUpdate {
	tu.mutation.AddAwayGameIDs(ids...)
	return tu
}

// AddAwayGames adds the "awayGames" edges to the Game entity.
func (tu *TeamUpdate) AddAwayGames(g ...*Game) *TeamUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.AddAwayGameIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tu *TeamUpdate) Mutation() *TeamMutation {
	return tu.mutation
}

// ClearPlayers clears all "players" edges to the Player entity.
func (tu *TeamUpdate) ClearPlayers() *TeamUpdate {
	tu.mutation.ClearPlayers()
	return tu
}

// RemovePlayerIDs removes the "players" edge to Player entities by IDs.
func (tu *TeamUpdate) RemovePlayerIDs(ids ...int) *TeamUpdate {
	tu.mutation.RemovePlayerIDs(ids...)
	return tu
}

// RemovePlayers removes "players" edges to Player entities.
func (tu *TeamUpdate) RemovePlayers(p ...*Player) *TeamUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.RemovePlayerIDs(ids...)
}

// ClearHomeGames clears all "homeGames" edges to the Game entity.
func (tu *TeamUpdate) ClearHomeGames() *TeamUpdate {
	tu.mutation.ClearHomeGames()
	return tu
}

// RemoveHomeGameIDs removes the "homeGames" edge to Game entities by IDs.
func (tu *TeamUpdate) RemoveHomeGameIDs(ids ...int) *TeamUpdate {
	tu.mutation.RemoveHomeGameIDs(ids...)
	return tu
}

// RemoveHomeGames removes "homeGames" edges to Game entities.
func (tu *TeamUpdate) RemoveHomeGames(g ...*Game) *TeamUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.RemoveHomeGameIDs(ids...)
}

// ClearAwayGames clears all "awayGames" edges to the Game entity.
func (tu *TeamUpdate) ClearAwayGames() *TeamUpdate {
	tu.mutation.ClearAwayGames()
	return tu
}

// RemoveAwayGameIDs removes the "awayGames" edge to Game entities by IDs.
func (tu *TeamUpdate) RemoveAwayGameIDs(ids ...int) *TeamUpdate {
	tu.mutation.RemoveAwayGameIDs(ids...)
	return tu
}

// RemoveAwayGames removes "awayGames" edges to Game entities.
func (tu *TeamUpdate) RemoveAwayGames(g ...*Game) *TeamUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.RemoveAwayGameIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeamUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeamUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeamUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeamUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TeamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.LogoFilepath(); ok {
		_spec.SetField(team.FieldLogoFilepath, field.TypeString, value)
	}
	if value, ok := tu.mutation.Eliminated(); ok {
		_spec.SetField(team.FieldEliminated, field.TypeBool, value)
	}
	if tu.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.PlayersTable,
			Columns: []string{team.PlayersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedPlayersIDs(); len(nodes) > 0 && !tu.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.PlayersTable,
			Columns: []string{team.PlayersColumn},
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
	if nodes := tu.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.PlayersTable,
			Columns: []string{team.PlayersColumn},
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
	if tu.mutation.HomeGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.HomeGamesTable,
			Columns: []string{team.HomeGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedHomeGamesIDs(); len(nodes) > 0 && !tu.mutation.HomeGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.HomeGamesTable,
			Columns: []string{team.HomeGamesColumn},
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
	if nodes := tu.mutation.HomeGamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.HomeGamesTable,
			Columns: []string{team.HomeGamesColumn},
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
	if tu.mutation.AwayGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.AwayGamesTable,
			Columns: []string{team.AwayGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedAwayGamesIDs(); len(nodes) > 0 && !tu.mutation.AwayGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.AwayGamesTable,
			Columns: []string{team.AwayGamesColumn},
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
	if nodes := tu.mutation.AwayGamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.AwayGamesTable,
			Columns: []string{team.AwayGamesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TeamUpdateOne is the builder for updating a single Team entity.
type TeamUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeamMutation
}

// SetName sets the "name" field.
func (tuo *TeamUpdateOne) SetName(s string) *TeamUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetLogoFilepath sets the "logo_filepath" field.
func (tuo *TeamUpdateOne) SetLogoFilepath(s string) *TeamUpdateOne {
	tuo.mutation.SetLogoFilepath(s)
	return tuo
}

// SetEliminated sets the "eliminated" field.
func (tuo *TeamUpdateOne) SetEliminated(b bool) *TeamUpdateOne {
	tuo.mutation.SetEliminated(b)
	return tuo
}

// SetNillableEliminated sets the "eliminated" field if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableEliminated(b *bool) *TeamUpdateOne {
	if b != nil {
		tuo.SetEliminated(*b)
	}
	return tuo
}

// AddPlayerIDs adds the "players" edge to the Player entity by IDs.
func (tuo *TeamUpdateOne) AddPlayerIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.AddPlayerIDs(ids...)
	return tuo
}

// AddPlayers adds the "players" edges to the Player entity.
func (tuo *TeamUpdateOne) AddPlayers(p ...*Player) *TeamUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.AddPlayerIDs(ids...)
}

// AddHomeGameIDs adds the "homeGames" edge to the Game entity by IDs.
func (tuo *TeamUpdateOne) AddHomeGameIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.AddHomeGameIDs(ids...)
	return tuo
}

// AddHomeGames adds the "homeGames" edges to the Game entity.
func (tuo *TeamUpdateOne) AddHomeGames(g ...*Game) *TeamUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.AddHomeGameIDs(ids...)
}

// AddAwayGameIDs adds the "awayGames" edge to the Game entity by IDs.
func (tuo *TeamUpdateOne) AddAwayGameIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.AddAwayGameIDs(ids...)
	return tuo
}

// AddAwayGames adds the "awayGames" edges to the Game entity.
func (tuo *TeamUpdateOne) AddAwayGames(g ...*Game) *TeamUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.AddAwayGameIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tuo *TeamUpdateOne) Mutation() *TeamMutation {
	return tuo.mutation
}

// ClearPlayers clears all "players" edges to the Player entity.
func (tuo *TeamUpdateOne) ClearPlayers() *TeamUpdateOne {
	tuo.mutation.ClearPlayers()
	return tuo
}

// RemovePlayerIDs removes the "players" edge to Player entities by IDs.
func (tuo *TeamUpdateOne) RemovePlayerIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.RemovePlayerIDs(ids...)
	return tuo
}

// RemovePlayers removes "players" edges to Player entities.
func (tuo *TeamUpdateOne) RemovePlayers(p ...*Player) *TeamUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.RemovePlayerIDs(ids...)
}

// ClearHomeGames clears all "homeGames" edges to the Game entity.
func (tuo *TeamUpdateOne) ClearHomeGames() *TeamUpdateOne {
	tuo.mutation.ClearHomeGames()
	return tuo
}

// RemoveHomeGameIDs removes the "homeGames" edge to Game entities by IDs.
func (tuo *TeamUpdateOne) RemoveHomeGameIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.RemoveHomeGameIDs(ids...)
	return tuo
}

// RemoveHomeGames removes "homeGames" edges to Game entities.
func (tuo *TeamUpdateOne) RemoveHomeGames(g ...*Game) *TeamUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.RemoveHomeGameIDs(ids...)
}

// ClearAwayGames clears all "awayGames" edges to the Game entity.
func (tuo *TeamUpdateOne) ClearAwayGames() *TeamUpdateOne {
	tuo.mutation.ClearAwayGames()
	return tuo
}

// RemoveAwayGameIDs removes the "awayGames" edge to Game entities by IDs.
func (tuo *TeamUpdateOne) RemoveAwayGameIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.RemoveAwayGameIDs(ids...)
	return tuo
}

// RemoveAwayGames removes "awayGames" edges to Game entities.
func (tuo *TeamUpdateOne) RemoveAwayGames(g ...*Game) *TeamUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.RemoveAwayGameIDs(ids...)
}

// Where appends a list predicates to the TeamUpdate builder.
func (tuo *TeamUpdateOne) Where(ps ...predicate.Team) *TeamUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeamUpdateOne) Select(field string, fields ...string) *TeamUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Team entity.
func (tuo *TeamUpdateOne) Save(ctx context.Context) (*Team, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeamUpdateOne) SaveX(ctx context.Context) *Team {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeamUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeamUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TeamUpdateOne) sqlSave(ctx context.Context) (_node *Team, err error) {
	_spec := sqlgraph.NewUpdateSpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Team.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, team.FieldID)
		for _, f := range fields {
			if !team.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != team.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.LogoFilepath(); ok {
		_spec.SetField(team.FieldLogoFilepath, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Eliminated(); ok {
		_spec.SetField(team.FieldEliminated, field.TypeBool, value)
	}
	if tuo.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.PlayersTable,
			Columns: []string{team.PlayersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedPlayersIDs(); len(nodes) > 0 && !tuo.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.PlayersTable,
			Columns: []string{team.PlayersColumn},
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
	if nodes := tuo.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.PlayersTable,
			Columns: []string{team.PlayersColumn},
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
	if tuo.mutation.HomeGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.HomeGamesTable,
			Columns: []string{team.HomeGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedHomeGamesIDs(); len(nodes) > 0 && !tuo.mutation.HomeGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.HomeGamesTable,
			Columns: []string{team.HomeGamesColumn},
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
	if nodes := tuo.mutation.HomeGamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.HomeGamesTable,
			Columns: []string{team.HomeGamesColumn},
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
	if tuo.mutation.AwayGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.AwayGamesTable,
			Columns: []string{team.AwayGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedAwayGamesIDs(); len(nodes) > 0 && !tuo.mutation.AwayGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.AwayGamesTable,
			Columns: []string{team.AwayGamesColumn},
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
	if nodes := tuo.mutation.AwayGamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.AwayGamesTable,
			Columns: []string{team.AwayGamesColumn},
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
	_node = &Team{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
