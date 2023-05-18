// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EntriesColumns holds the columns for the "entries" table.
	EntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "owner_name", Type: field.TypeString},
		{Name: "league_entries", Type: field.TypeInt, Nullable: true},
	}
	// EntriesTable holds the schema information for the "entries" table.
	EntriesTable = &schema.Table{
		Name:       "entries",
		Columns:    EntriesColumns,
		PrimaryKey: []*schema.Column{EntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "entries_leagues_entries",
				Columns:    []*schema.Column{EntriesColumns[2]},
				RefColumns: []*schema.Column{LeaguesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GamesColumns holds the columns for the "games" table.
	GamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "home_score", Type: field.TypeInt},
		{Name: "away_score", Type: field.TypeInt},
		{Name: "game_stats_game", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "team_home_games", Type: field.TypeInt},
		{Name: "team_away_games", Type: field.TypeInt},
	}
	// GamesTable holds the schema information for the "games" table.
	GamesTable = &schema.Table{
		Name:       "games",
		Columns:    GamesColumns,
		PrimaryKey: []*schema.Column{GamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "games_game_stats_game",
				Columns:    []*schema.Column{GamesColumns[3]},
				RefColumns: []*schema.Column{GameStatsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "games_teams_homeGames",
				Columns:    []*schema.Column{GamesColumns[4]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "games_teams_awayGames",
				Columns:    []*schema.Column{GamesColumns[5]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GameStatsColumns holds the columns for the "game_stats" table.
	GameStatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "goals", Type: field.TypeInt, Default: 0},
		{Name: "assists", Type: field.TypeInt, Default: 0},
		{Name: "win", Type: field.TypeBool, Default: false},
		{Name: "shutout", Type: field.TypeBool, Default: false},
		{Name: "home_game", Type: field.TypeBool, Default: false},
		{Name: "game_stats_player", Type: field.TypeInt},
	}
	// GameStatsTable holds the schema information for the "game_stats" table.
	GameStatsTable = &schema.Table{
		Name:       "game_stats",
		Columns:    GameStatsColumns,
		PrimaryKey: []*schema.Column{GameStatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "game_stats_players_player",
				Columns:    []*schema.Column{GameStatsColumns[6]},
				RefColumns: []*schema.Column{PlayersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LeaguesColumns holds the columns for the "leagues" table.
	LeaguesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "season", Type: field.TypeInt},
		{Name: "public", Type: field.TypeBool},
		{Name: "num_forwards", Type: field.TypeInt},
		{Name: "num_defenders", Type: field.TypeInt},
		{Name: "num_goalies", Type: field.TypeInt},
	}
	// LeaguesTable holds the schema information for the "leagues" table.
	LeaguesTable = &schema.Table{
		Name:       "leagues",
		Columns:    LeaguesColumns,
		PrimaryKey: []*schema.Column{LeaguesColumns[0]},
	}
	// PlayersColumns holds the columns for the "players" table.
	PlayersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "position", Type: field.TypeEnum, Enums: []string{"Forward", "Defenseman", "Goalie"}},
		{Name: "stats_player", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "team_players", Type: field.TypeInt, Nullable: true},
	}
	// PlayersTable holds the schema information for the "players" table.
	PlayersTable = &schema.Table{
		Name:       "players",
		Columns:    PlayersColumns,
		PrimaryKey: []*schema.Column{PlayersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "players_stats_player",
				Columns:    []*schema.Column{PlayersColumns[3]},
				RefColumns: []*schema.Column{StatsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "players_teams_players",
				Columns:    []*schema.Column{PlayersColumns[4]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StatsColumns holds the columns for the "stats" table.
	StatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "goals", Type: field.TypeInt, Default: 0},
		{Name: "assists", Type: field.TypeInt, Default: 0},
		{Name: "shutouts", Type: field.TypeInt, Default: 0},
		{Name: "wins", Type: field.TypeInt, Default: 0},
	}
	// StatsTable holds the schema information for the "stats" table.
	StatsTable = &schema.Table{
		Name:       "stats",
		Columns:    StatsColumns,
		PrimaryKey: []*schema.Column{StatsColumns[0]},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "logo_filepath", Type: field.TypeString, Unique: true},
		{Name: "eliminated", Type: field.TypeBool, Default: false},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// EntryPlayersColumns holds the columns for the "entry_players" table.
	EntryPlayersColumns = []*schema.Column{
		{Name: "entry_id", Type: field.TypeInt},
		{Name: "player_id", Type: field.TypeInt},
	}
	// EntryPlayersTable holds the schema information for the "entry_players" table.
	EntryPlayersTable = &schema.Table{
		Name:       "entry_players",
		Columns:    EntryPlayersColumns,
		PrimaryKey: []*schema.Column{EntryPlayersColumns[0], EntryPlayersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "entry_players_entry_id",
				Columns:    []*schema.Column{EntryPlayersColumns[0]},
				RefColumns: []*schema.Column{EntriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "entry_players_player_id",
				Columns:    []*schema.Column{EntryPlayersColumns[1]},
				RefColumns: []*schema.Column{PlayersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EntriesTable,
		GamesTable,
		GameStatsTable,
		LeaguesTable,
		PlayersTable,
		StatsTable,
		TeamsTable,
		EntryPlayersTable,
	}
)

func init() {
	EntriesTable.ForeignKeys[0].RefTable = LeaguesTable
	GamesTable.ForeignKeys[0].RefTable = GameStatsTable
	GamesTable.ForeignKeys[1].RefTable = TeamsTable
	GamesTable.ForeignKeys[2].RefTable = TeamsTable
	GameStatsTable.ForeignKeys[0].RefTable = PlayersTable
	PlayersTable.ForeignKeys[0].RefTable = StatsTable
	PlayersTable.ForeignKeys[1].RefTable = TeamsTable
	EntryPlayersTable.ForeignKeys[0].RefTable = EntriesTable
	EntryPlayersTable.ForeignKeys[1].RefTable = PlayersTable
}
