// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/game"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/goaliestats"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/predicate"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/skaterstats"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/team"
)

// GameQuery is the builder for querying Game entities.
type GameQuery struct {
	config
	ctx             *QueryContext
	order           []game.OrderOption
	inters          []Interceptor
	predicates      []predicate.Game
	withAwayTeam    *TeamQuery
	withHomeTeam    *TeamQuery
	withSkaterStats *SkaterStatsQuery
	withGoalieStats *GoalieStatsQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GameQuery builder.
func (gq *GameQuery) Where(ps ...predicate.Game) *GameQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit the number of records to be returned by this query.
func (gq *GameQuery) Limit(limit int) *GameQuery {
	gq.ctx.Limit = &limit
	return gq
}

// Offset to start from.
func (gq *GameQuery) Offset(offset int) *GameQuery {
	gq.ctx.Offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GameQuery) Unique(unique bool) *GameQuery {
	gq.ctx.Unique = &unique
	return gq
}

// Order specifies how the records should be ordered.
func (gq *GameQuery) Order(o ...game.OrderOption) *GameQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryAwayTeam chains the current query on the "awayTeam" edge.
func (gq *GameQuery) QueryAwayTeam() *TeamQuery {
	query := (&TeamClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, game.AwayTeamTable, game.AwayTeamColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHomeTeam chains the current query on the "homeTeam" edge.
func (gq *GameQuery) QueryHomeTeam() *TeamQuery {
	query := (&TeamClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, game.HomeTeamTable, game.HomeTeamColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySkaterStats chains the current query on the "skaterStats" edge.
func (gq *GameQuery) QuerySkaterStats() *SkaterStatsQuery {
	query := (&SkaterStatsClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, selector),
			sqlgraph.To(skaterstats.Table, skaterstats.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, game.SkaterStatsTable, game.SkaterStatsColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGoalieStats chains the current query on the "goalieStats" edge.
func (gq *GameQuery) QueryGoalieStats() *GoalieStatsQuery {
	query := (&GoalieStatsClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, selector),
			sqlgraph.To(goaliestats.Table, goaliestats.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, game.GoalieStatsTable, game.GoalieStatsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Game entity from the query.
// Returns a *NotFoundError when no Game was found.
func (gq *GameQuery) First(ctx context.Context) (*Game, error) {
	nodes, err := gq.Limit(1).All(setContextOp(ctx, gq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{game.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GameQuery) FirstX(ctx context.Context) *Game {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Game ID from the query.
// Returns a *NotFoundError when no Game ID was found.
func (gq *GameQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(setContextOp(ctx, gq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{game.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GameQuery) FirstIDX(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Game entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Game entity is found.
// Returns a *NotFoundError when no Game entities are found.
func (gq *GameQuery) Only(ctx context.Context) (*Game, error) {
	nodes, err := gq.Limit(2).All(setContextOp(ctx, gq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{game.Label}
	default:
		return nil, &NotSingularError{game.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GameQuery) OnlyX(ctx context.Context) *Game {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Game ID in the query.
// Returns a *NotSingularError when more than one Game ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GameQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(setContextOp(ctx, gq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{game.Label}
	default:
		err = &NotSingularError{game.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GameQuery) OnlyIDX(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Games.
func (gq *GameQuery) All(ctx context.Context) ([]*Game, error) {
	ctx = setContextOp(ctx, gq.ctx, "All")
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Game, *GameQuery]()
	return withInterceptors[[]*Game](ctx, gq, qr, gq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gq *GameQuery) AllX(ctx context.Context) []*Game {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Game IDs.
func (gq *GameQuery) IDs(ctx context.Context) (ids []int, err error) {
	if gq.ctx.Unique == nil && gq.path != nil {
		gq.Unique(true)
	}
	ctx = setContextOp(ctx, gq.ctx, "IDs")
	if err = gq.Select(game.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GameQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GameQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gq.ctx, "Count")
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gq, querierCount[*GameQuery](), gq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GameQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GameQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gq.ctx, "Exist")
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GameQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GameQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GameQuery) Clone() *GameQuery {
	if gq == nil {
		return nil
	}
	return &GameQuery{
		config:          gq.config,
		ctx:             gq.ctx.Clone(),
		order:           append([]game.OrderOption{}, gq.order...),
		inters:          append([]Interceptor{}, gq.inters...),
		predicates:      append([]predicate.Game{}, gq.predicates...),
		withAwayTeam:    gq.withAwayTeam.Clone(),
		withHomeTeam:    gq.withHomeTeam.Clone(),
		withSkaterStats: gq.withSkaterStats.Clone(),
		withGoalieStats: gq.withGoalieStats.Clone(),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// WithAwayTeam tells the query-builder to eager-load the nodes that are connected to
// the "awayTeam" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GameQuery) WithAwayTeam(opts ...func(*TeamQuery)) *GameQuery {
	query := (&TeamClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withAwayTeam = query
	return gq
}

// WithHomeTeam tells the query-builder to eager-load the nodes that are connected to
// the "homeTeam" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GameQuery) WithHomeTeam(opts ...func(*TeamQuery)) *GameQuery {
	query := (&TeamClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withHomeTeam = query
	return gq
}

// WithSkaterStats tells the query-builder to eager-load the nodes that are connected to
// the "skaterStats" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GameQuery) WithSkaterStats(opts ...func(*SkaterStatsQuery)) *GameQuery {
	query := (&SkaterStatsClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withSkaterStats = query
	return gq
}

// WithGoalieStats tells the query-builder to eager-load the nodes that are connected to
// the "goalieStats" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GameQuery) WithGoalieStats(opts ...func(*GoalieStatsQuery)) *GameQuery {
	query := (&GoalieStatsClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withGoalieStats = query
	return gq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HomeScore int `json:"homeScore,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Game.Query().
//		GroupBy(game.FieldHomeScore).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gq *GameQuery) GroupBy(field string, fields ...string) *GameGroupBy {
	gq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GameGroupBy{build: gq}
	grbuild.flds = &gq.ctx.Fields
	grbuild.label = game.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HomeScore int `json:"homeScore,omitempty"`
//	}
//
//	client.Game.Query().
//		Select(game.FieldHomeScore).
//		Scan(ctx, &v)
func (gq *GameQuery) Select(fields ...string) *GameSelect {
	gq.ctx.Fields = append(gq.ctx.Fields, fields...)
	sbuild := &GameSelect{GameQuery: gq}
	sbuild.label = game.Label
	sbuild.flds, sbuild.scan = &gq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GameSelect configured with the given aggregations.
func (gq *GameQuery) Aggregate(fns ...AggregateFunc) *GameSelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GameQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gq); err != nil {
				return err
			}
		}
	}
	for _, f := range gq.ctx.Fields {
		if !game.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GameQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Game, error) {
	var (
		nodes       = []*Game{}
		withFKs     = gq.withFKs
		_spec       = gq.querySpec()
		loadedTypes = [4]bool{
			gq.withAwayTeam != nil,
			gq.withHomeTeam != nil,
			gq.withSkaterStats != nil,
			gq.withGoalieStats != nil,
		}
	)
	if gq.withAwayTeam != nil || gq.withHomeTeam != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, game.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Game).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Game{config: gq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gq.withAwayTeam; query != nil {
		if err := gq.loadAwayTeam(ctx, query, nodes, nil,
			func(n *Game, e *Team) { n.Edges.AwayTeam = e }); err != nil {
			return nil, err
		}
	}
	if query := gq.withHomeTeam; query != nil {
		if err := gq.loadHomeTeam(ctx, query, nodes, nil,
			func(n *Game, e *Team) { n.Edges.HomeTeam = e }); err != nil {
			return nil, err
		}
	}
	if query := gq.withSkaterStats; query != nil {
		if err := gq.loadSkaterStats(ctx, query, nodes,
			func(n *Game) { n.Edges.SkaterStats = []*SkaterStats{} },
			func(n *Game, e *SkaterStats) { n.Edges.SkaterStats = append(n.Edges.SkaterStats, e) }); err != nil {
			return nil, err
		}
	}
	if query := gq.withGoalieStats; query != nil {
		if err := gq.loadGoalieStats(ctx, query, nodes,
			func(n *Game) { n.Edges.GoalieStats = []*GoalieStats{} },
			func(n *Game, e *GoalieStats) { n.Edges.GoalieStats = append(n.Edges.GoalieStats, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gq *GameQuery) loadAwayTeam(ctx context.Context, query *TeamQuery, nodes []*Game, init func(*Game), assign func(*Game, *Team)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Game)
	for i := range nodes {
		if nodes[i].team_away_games == nil {
			continue
		}
		fk := *nodes[i].team_away_games
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(team.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "team_away_games" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gq *GameQuery) loadHomeTeam(ctx context.Context, query *TeamQuery, nodes []*Game, init func(*Game), assign func(*Game, *Team)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Game)
	for i := range nodes {
		if nodes[i].team_home_games == nil {
			continue
		}
		fk := *nodes[i].team_home_games
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(team.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "team_home_games" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gq *GameQuery) loadSkaterStats(ctx context.Context, query *SkaterStatsQuery, nodes []*Game, init func(*Game), assign func(*Game, *SkaterStats)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Game)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.SkaterStats(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(game.SkaterStatsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.game_skater_stats
		if fk == nil {
			return fmt.Errorf(`foreign-key "game_skater_stats" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "game_skater_stats" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (gq *GameQuery) loadGoalieStats(ctx context.Context, query *GoalieStatsQuery, nodes []*Game, init func(*Game), assign func(*Game, *GoalieStats)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Game)
	nids := make(map[int]map[*Game]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(game.GoalieStatsTable)
		s.Join(joinT).On(s.C(goaliestats.FieldID), joinT.C(game.GoalieStatsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(game.GoalieStatsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(game.GoalieStatsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Game]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*GoalieStats](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "goalieStats" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (gq *GameQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	_spec.Node.Columns = gq.ctx.Fields
	if len(gq.ctx.Fields) > 0 {
		_spec.Unique = gq.ctx.Unique != nil && *gq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GameQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(game.Table, game.Columns, sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt))
	_spec.From = gq.sql
	if unique := gq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gq.path != nil {
		_spec.Unique = true
	}
	if fields := gq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, game.FieldID)
		for i := range fields {
			if fields[i] != game.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GameQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(game.Table)
	columns := gq.ctx.Fields
	if len(columns) == 0 {
		columns = game.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.ctx.Unique != nil && *gq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GameGroupBy is the group-by builder for Game entities.
type GameGroupBy struct {
	selector
	build *GameQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GameGroupBy) Aggregate(fns ...AggregateFunc) *GameGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GameGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, "GroupBy")
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GameQuery, *GameGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GameGroupBy) sqlScan(ctx context.Context, root *GameQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GameSelect is the builder for selecting fields of Game entities.
type GameSelect struct {
	*GameQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GameSelect) Aggregate(fns ...AggregateFunc) *GameSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GameSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, "Select")
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GameQuery, *GameSelect](ctx, gs.GameQuery, gs, gs.inters, v)
}

func (gs *GameSelect) sqlScan(ctx context.Context, root *GameQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
