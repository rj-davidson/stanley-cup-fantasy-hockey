// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/entry"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/game"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/league"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/player"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/team"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Entry is the client for interacting with the Entry builders.
	Entry *EntryClient
	// Game is the client for interacting with the Game builders.
	Game *GameClient
	// League is the client for interacting with the League builders.
	League *LeagueClient
	// Player is the client for interacting with the Player builders.
	Player *PlayerClient
	// Team is the client for interacting with the Team builders.
	Team *TeamClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Entry = NewEntryClient(c.config)
	c.Game = NewGameClient(c.config)
	c.League = NewLeagueClient(c.config)
	c.Player = NewPlayerClient(c.config)
	c.Team = NewTeamClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Entry:  NewEntryClient(cfg),
		Game:   NewGameClient(cfg),
		League: NewLeagueClient(cfg),
		Player: NewPlayerClient(cfg),
		Team:   NewTeamClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Entry:  NewEntryClient(cfg),
		Game:   NewGameClient(cfg),
		League: NewLeagueClient(cfg),
		Player: NewPlayerClient(cfg),
		Team:   NewTeamClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Entry.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Entry.Use(hooks...)
	c.Game.Use(hooks...)
	c.League.Use(hooks...)
	c.Player.Use(hooks...)
	c.Team.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Entry.Intercept(interceptors...)
	c.Game.Intercept(interceptors...)
	c.League.Intercept(interceptors...)
	c.Player.Intercept(interceptors...)
	c.Team.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *EntryMutation:
		return c.Entry.mutate(ctx, m)
	case *GameMutation:
		return c.Game.mutate(ctx, m)
	case *LeagueMutation:
		return c.League.mutate(ctx, m)
	case *PlayerMutation:
		return c.Player.mutate(ctx, m)
	case *TeamMutation:
		return c.Team.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// EntryClient is a client for the Entry schema.
type EntryClient struct {
	config
}

// NewEntryClient returns a client for the Entry from the given config.
func NewEntryClient(c config) *EntryClient {
	return &EntryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `entry.Hooks(f(g(h())))`.
func (c *EntryClient) Use(hooks ...Hook) {
	c.hooks.Entry = append(c.hooks.Entry, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `entry.Intercept(f(g(h())))`.
func (c *EntryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Entry = append(c.inters.Entry, interceptors...)
}

// Create returns a builder for creating a Entry entity.
func (c *EntryClient) Create() *EntryCreate {
	mutation := newEntryMutation(c.config, OpCreate)
	return &EntryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Entry entities.
func (c *EntryClient) CreateBulk(builders ...*EntryCreate) *EntryCreateBulk {
	return &EntryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Entry.
func (c *EntryClient) Update() *EntryUpdate {
	mutation := newEntryMutation(c.config, OpUpdate)
	return &EntryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EntryClient) UpdateOne(e *Entry) *EntryUpdateOne {
	mutation := newEntryMutation(c.config, OpUpdateOne, withEntry(e))
	return &EntryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EntryClient) UpdateOneID(id int) *EntryUpdateOne {
	mutation := newEntryMutation(c.config, OpUpdateOne, withEntryID(id))
	return &EntryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Entry.
func (c *EntryClient) Delete() *EntryDelete {
	mutation := newEntryMutation(c.config, OpDelete)
	return &EntryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EntryClient) DeleteOne(e *Entry) *EntryDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EntryClient) DeleteOneID(id int) *EntryDeleteOne {
	builder := c.Delete().Where(entry.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EntryDeleteOne{builder}
}

// Query returns a query builder for Entry.
func (c *EntryClient) Query() *EntryQuery {
	return &EntryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEntry},
		inters: c.Interceptors(),
	}
}

// Get returns a Entry entity by its id.
func (c *EntryClient) Get(ctx context.Context, id int) (*Entry, error) {
	return c.Query().Where(entry.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EntryClient) GetX(ctx context.Context, id int) *Entry {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLeague queries the league edge of a Entry.
func (c *EntryClient) QueryLeague(e *Entry) *LeagueQuery {
	query := (&LeagueClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(entry.Table, entry.FieldID, id),
			sqlgraph.To(league.Table, league.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entry.LeagueTable, entry.LeagueColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPlayers queries the players edge of a Entry.
func (c *EntryClient) QueryPlayers(e *Entry) *PlayerQuery {
	query := (&PlayerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(entry.Table, entry.FieldID, id),
			sqlgraph.To(player.Table, player.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, entry.PlayersTable, entry.PlayersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EntryClient) Hooks() []Hook {
	return c.hooks.Entry
}

// Interceptors returns the client interceptors.
func (c *EntryClient) Interceptors() []Interceptor {
	return c.inters.Entry
}

func (c *EntryClient) mutate(ctx context.Context, m *EntryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EntryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EntryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EntryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EntryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Entry mutation op: %q", m.Op())
	}
}

// GameClient is a client for the Game schema.
type GameClient struct {
	config
}

// NewGameClient returns a client for the Game from the given config.
func NewGameClient(c config) *GameClient {
	return &GameClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `game.Hooks(f(g(h())))`.
func (c *GameClient) Use(hooks ...Hook) {
	c.hooks.Game = append(c.hooks.Game, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `game.Intercept(f(g(h())))`.
func (c *GameClient) Intercept(interceptors ...Interceptor) {
	c.inters.Game = append(c.inters.Game, interceptors...)
}

// Create returns a builder for creating a Game entity.
func (c *GameClient) Create() *GameCreate {
	mutation := newGameMutation(c.config, OpCreate)
	return &GameCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Game entities.
func (c *GameClient) CreateBulk(builders ...*GameCreate) *GameCreateBulk {
	return &GameCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Game.
func (c *GameClient) Update() *GameUpdate {
	mutation := newGameMutation(c.config, OpUpdate)
	return &GameUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GameClient) UpdateOne(ga *Game) *GameUpdateOne {
	mutation := newGameMutation(c.config, OpUpdateOne, withGame(ga))
	return &GameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GameClient) UpdateOneID(id int) *GameUpdateOne {
	mutation := newGameMutation(c.config, OpUpdateOne, withGameID(id))
	return &GameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Game.
func (c *GameClient) Delete() *GameDelete {
	mutation := newGameMutation(c.config, OpDelete)
	return &GameDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GameClient) DeleteOne(ga *Game) *GameDeleteOne {
	return c.DeleteOneID(ga.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GameClient) DeleteOneID(id int) *GameDeleteOne {
	builder := c.Delete().Where(game.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GameDeleteOne{builder}
}

// Query returns a query builder for Game.
func (c *GameClient) Query() *GameQuery {
	return &GameQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGame},
		inters: c.Interceptors(),
	}
}

// Get returns a Game entity by its id.
func (c *GameClient) Get(ctx context.Context, id int) (*Game, error) {
	return c.Query().Where(game.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GameClient) GetX(ctx context.Context, id int) *Game {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAwayTeam queries the awayTeam edge of a Game.
func (c *GameClient) QueryAwayTeam(ga *Game) *TeamQuery {
	query := (&TeamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, game.AwayTeamTable, game.AwayTeamColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHomeTeam queries the homeTeam edge of a Game.
func (c *GameClient) QueryHomeTeam(ga *Game) *TeamQuery {
	query := (&TeamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, game.HomeTeamTable, game.HomeTeamColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAwayGoalie queries the awayGoalie edge of a Game.
func (c *GameClient) QueryAwayGoalie(ga *Game) *PlayerQuery {
	query := (&PlayerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, id),
			sqlgraph.To(player.Table, player.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, game.AwayGoalieTable, game.AwayGoalieColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHomeGoalie queries the homeGoalie edge of a Game.
func (c *GameClient) QueryHomeGoalie(ga *Game) *PlayerQuery {
	query := (&PlayerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, id),
			sqlgraph.To(player.Table, player.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, game.HomeGoalieTable, game.HomeGoalieColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GameClient) Hooks() []Hook {
	return c.hooks.Game
}

// Interceptors returns the client interceptors.
func (c *GameClient) Interceptors() []Interceptor {
	return c.inters.Game
}

func (c *GameClient) mutate(ctx context.Context, m *GameMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GameCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GameUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GameDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Game mutation op: %q", m.Op())
	}
}

// LeagueClient is a client for the League schema.
type LeagueClient struct {
	config
}

// NewLeagueClient returns a client for the League from the given config.
func NewLeagueClient(c config) *LeagueClient {
	return &LeagueClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `league.Hooks(f(g(h())))`.
func (c *LeagueClient) Use(hooks ...Hook) {
	c.hooks.League = append(c.hooks.League, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `league.Intercept(f(g(h())))`.
func (c *LeagueClient) Intercept(interceptors ...Interceptor) {
	c.inters.League = append(c.inters.League, interceptors...)
}

// Create returns a builder for creating a League entity.
func (c *LeagueClient) Create() *LeagueCreate {
	mutation := newLeagueMutation(c.config, OpCreate)
	return &LeagueCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of League entities.
func (c *LeagueClient) CreateBulk(builders ...*LeagueCreate) *LeagueCreateBulk {
	return &LeagueCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for League.
func (c *LeagueClient) Update() *LeagueUpdate {
	mutation := newLeagueMutation(c.config, OpUpdate)
	return &LeagueUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LeagueClient) UpdateOne(l *League) *LeagueUpdateOne {
	mutation := newLeagueMutation(c.config, OpUpdateOne, withLeague(l))
	return &LeagueUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LeagueClient) UpdateOneID(id int) *LeagueUpdateOne {
	mutation := newLeagueMutation(c.config, OpUpdateOne, withLeagueID(id))
	return &LeagueUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for League.
func (c *LeagueClient) Delete() *LeagueDelete {
	mutation := newLeagueMutation(c.config, OpDelete)
	return &LeagueDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LeagueClient) DeleteOne(l *League) *LeagueDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LeagueClient) DeleteOneID(id int) *LeagueDeleteOne {
	builder := c.Delete().Where(league.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LeagueDeleteOne{builder}
}

// Query returns a query builder for League.
func (c *LeagueClient) Query() *LeagueQuery {
	return &LeagueQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLeague},
		inters: c.Interceptors(),
	}
}

// Get returns a League entity by its id.
func (c *LeagueClient) Get(ctx context.Context, id int) (*League, error) {
	return c.Query().Where(league.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LeagueClient) GetX(ctx context.Context, id int) *League {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEntries queries the entries edge of a League.
func (c *LeagueClient) QueryEntries(l *League) *EntryQuery {
	query := (&EntryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(league.Table, league.FieldID, id),
			sqlgraph.To(entry.Table, entry.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, league.EntriesTable, league.EntriesColumn),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LeagueClient) Hooks() []Hook {
	return c.hooks.League
}

// Interceptors returns the client interceptors.
func (c *LeagueClient) Interceptors() []Interceptor {
	return c.inters.League
}

func (c *LeagueClient) mutate(ctx context.Context, m *LeagueMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LeagueCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LeagueUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LeagueUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LeagueDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown League mutation op: %q", m.Op())
	}
}

// PlayerClient is a client for the Player schema.
type PlayerClient struct {
	config
}

// NewPlayerClient returns a client for the Player from the given config.
func NewPlayerClient(c config) *PlayerClient {
	return &PlayerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `player.Hooks(f(g(h())))`.
func (c *PlayerClient) Use(hooks ...Hook) {
	c.hooks.Player = append(c.hooks.Player, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `player.Intercept(f(g(h())))`.
func (c *PlayerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Player = append(c.inters.Player, interceptors...)
}

// Create returns a builder for creating a Player entity.
func (c *PlayerClient) Create() *PlayerCreate {
	mutation := newPlayerMutation(c.config, OpCreate)
	return &PlayerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Player entities.
func (c *PlayerClient) CreateBulk(builders ...*PlayerCreate) *PlayerCreateBulk {
	return &PlayerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Player.
func (c *PlayerClient) Update() *PlayerUpdate {
	mutation := newPlayerMutation(c.config, OpUpdate)
	return &PlayerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlayerClient) UpdateOne(pl *Player) *PlayerUpdateOne {
	mutation := newPlayerMutation(c.config, OpUpdateOne, withPlayer(pl))
	return &PlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlayerClient) UpdateOneID(id int) *PlayerUpdateOne {
	mutation := newPlayerMutation(c.config, OpUpdateOne, withPlayerID(id))
	return &PlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Player.
func (c *PlayerClient) Delete() *PlayerDelete {
	mutation := newPlayerMutation(c.config, OpDelete)
	return &PlayerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PlayerClient) DeleteOne(pl *Player) *PlayerDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PlayerClient) DeleteOneID(id int) *PlayerDeleteOne {
	builder := c.Delete().Where(player.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlayerDeleteOne{builder}
}

// Query returns a query builder for Player.
func (c *PlayerClient) Query() *PlayerQuery {
	return &PlayerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePlayer},
		inters: c.Interceptors(),
	}
}

// Get returns a Player entity by its id.
func (c *PlayerClient) Get(ctx context.Context, id int) (*Player, error) {
	return c.Query().Where(player.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlayerClient) GetX(ctx context.Context, id int) *Player {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTeam queries the team edge of a Player.
func (c *PlayerClient) QueryTeam(pl *Player) *TeamQuery {
	query := (&TeamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(player.Table, player.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, player.TeamTable, player.TeamColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEntries queries the entries edge of a Player.
func (c *PlayerClient) QueryEntries(pl *Player) *EntryQuery {
	query := (&EntryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(player.Table, player.FieldID, id),
			sqlgraph.To(entry.Table, entry.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, player.EntriesTable, player.EntriesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHomeGamesAsGoalie queries the homeGamesAsGoalie edge of a Player.
func (c *PlayerClient) QueryHomeGamesAsGoalie(pl *Player) *GameQuery {
	query := (&GameClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(player.Table, player.FieldID, id),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, player.HomeGamesAsGoalieTable, player.HomeGamesAsGoalieColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAwayGamesAsGoalie queries the awayGamesAsGoalie edge of a Player.
func (c *PlayerClient) QueryAwayGamesAsGoalie(pl *Player) *GameQuery {
	query := (&GameClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(player.Table, player.FieldID, id),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, player.AwayGamesAsGoalieTable, player.AwayGamesAsGoalieColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PlayerClient) Hooks() []Hook {
	return c.hooks.Player
}

// Interceptors returns the client interceptors.
func (c *PlayerClient) Interceptors() []Interceptor {
	return c.inters.Player
}

func (c *PlayerClient) mutate(ctx context.Context, m *PlayerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PlayerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PlayerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PlayerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Player mutation op: %q", m.Op())
	}
}

// TeamClient is a client for the Team schema.
type TeamClient struct {
	config
}

// NewTeamClient returns a client for the Team from the given config.
func NewTeamClient(c config) *TeamClient {
	return &TeamClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `team.Hooks(f(g(h())))`.
func (c *TeamClient) Use(hooks ...Hook) {
	c.hooks.Team = append(c.hooks.Team, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `team.Intercept(f(g(h())))`.
func (c *TeamClient) Intercept(interceptors ...Interceptor) {
	c.inters.Team = append(c.inters.Team, interceptors...)
}

// Create returns a builder for creating a Team entity.
func (c *TeamClient) Create() *TeamCreate {
	mutation := newTeamMutation(c.config, OpCreate)
	return &TeamCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Team entities.
func (c *TeamClient) CreateBulk(builders ...*TeamCreate) *TeamCreateBulk {
	return &TeamCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Team.
func (c *TeamClient) Update() *TeamUpdate {
	mutation := newTeamMutation(c.config, OpUpdate)
	return &TeamUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TeamClient) UpdateOne(t *Team) *TeamUpdateOne {
	mutation := newTeamMutation(c.config, OpUpdateOne, withTeam(t))
	return &TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TeamClient) UpdateOneID(id int) *TeamUpdateOne {
	mutation := newTeamMutation(c.config, OpUpdateOne, withTeamID(id))
	return &TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Team.
func (c *TeamClient) Delete() *TeamDelete {
	mutation := newTeamMutation(c.config, OpDelete)
	return &TeamDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TeamClient) DeleteOne(t *Team) *TeamDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TeamClient) DeleteOneID(id int) *TeamDeleteOne {
	builder := c.Delete().Where(team.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TeamDeleteOne{builder}
}

// Query returns a query builder for Team.
func (c *TeamClient) Query() *TeamQuery {
	return &TeamQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTeam},
		inters: c.Interceptors(),
	}
}

// Get returns a Team entity by its id.
func (c *TeamClient) Get(ctx context.Context, id int) (*Team, error) {
	return c.Query().Where(team.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TeamClient) GetX(ctx context.Context, id int) *Team {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlayers queries the players edge of a Team.
func (c *TeamClient) QueryPlayers(t *Team) *PlayerQuery {
	query := (&PlayerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(player.Table, player.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, team.PlayersTable, team.PlayersColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHomeGames queries the homeGames edge of a Team.
func (c *TeamClient) QueryHomeGames(t *Team) *GameQuery {
	query := (&GameClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, team.HomeGamesTable, team.HomeGamesColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAwayGames queries the awayGames edge of a Team.
func (c *TeamClient) QueryAwayGames(t *Team) *GameQuery {
	query := (&GameClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, team.AwayGamesTable, team.AwayGamesColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TeamClient) Hooks() []Hook {
	return c.hooks.Team
}

// Interceptors returns the client interceptors.
func (c *TeamClient) Interceptors() []Interceptor {
	return c.inters.Team
}

func (c *TeamClient) mutate(ctx context.Context, m *TeamMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TeamCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TeamUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TeamDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TeamID mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Entry, Game, League, Player, Team []ent.Hook
	}
	inters struct {
		Entry, Game, League, Player, Team []ent.Interceptor
	}
)
