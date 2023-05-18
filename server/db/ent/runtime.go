// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/game"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/gamestats"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/player"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/schema"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/stats"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/team"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	gameFields := schema.Game{}.Fields()
	_ = gameFields
	// gameDescID is the schema descriptor for id field.
	gameDescID := gameFields[0].Descriptor()
	// game.IDValidator is a validator for the "id" field. It is called by the builders before save.
	game.IDValidator = gameDescID.Validators[0].(func(int) error)
	gamestatsFields := schema.GameStats{}.Fields()
	_ = gamestatsFields
	// gamestatsDescGoals is the schema descriptor for goals field.
	gamestatsDescGoals := gamestatsFields[0].Descriptor()
	// gamestats.DefaultGoals holds the default value on creation for the goals field.
	gamestats.DefaultGoals = gamestatsDescGoals.Default.(int)
	// gamestatsDescAssists is the schema descriptor for assists field.
	gamestatsDescAssists := gamestatsFields[1].Descriptor()
	// gamestats.DefaultAssists holds the default value on creation for the assists field.
	gamestats.DefaultAssists = gamestatsDescAssists.Default.(int)
	// gamestatsDescWin is the schema descriptor for win field.
	gamestatsDescWin := gamestatsFields[2].Descriptor()
	// gamestats.DefaultWin holds the default value on creation for the win field.
	gamestats.DefaultWin = gamestatsDescWin.Default.(bool)
	// gamestatsDescShutout is the schema descriptor for shutout field.
	gamestatsDescShutout := gamestatsFields[3].Descriptor()
	// gamestats.DefaultShutout holds the default value on creation for the shutout field.
	gamestats.DefaultShutout = gamestatsDescShutout.Default.(bool)
	// gamestatsDescHomeGame is the schema descriptor for homeGame field.
	gamestatsDescHomeGame := gamestatsFields[4].Descriptor()
	// gamestats.DefaultHomeGame holds the default value on creation for the homeGame field.
	gamestats.DefaultHomeGame = gamestatsDescHomeGame.Default.(bool)
	playerFields := schema.Player{}.Fields()
	_ = playerFields
	// playerDescID is the schema descriptor for id field.
	playerDescID := playerFields[0].Descriptor()
	// player.IDValidator is a validator for the "id" field. It is called by the builders before save.
	player.IDValidator = playerDescID.Validators[0].(func(int) error)
	statsFields := schema.Stats{}.Fields()
	_ = statsFields
	// statsDescGoals is the schema descriptor for goals field.
	statsDescGoals := statsFields[0].Descriptor()
	// stats.DefaultGoals holds the default value on creation for the goals field.
	stats.DefaultGoals = statsDescGoals.Default.(int)
	// statsDescAssists is the schema descriptor for assists field.
	statsDescAssists := statsFields[1].Descriptor()
	// stats.DefaultAssists holds the default value on creation for the assists field.
	stats.DefaultAssists = statsDescAssists.Default.(int)
	// statsDescShutouts is the schema descriptor for shutouts field.
	statsDescShutouts := statsFields[2].Descriptor()
	// stats.DefaultShutouts holds the default value on creation for the shutouts field.
	stats.DefaultShutouts = statsDescShutouts.Default.(int)
	// statsDescWins is the schema descriptor for wins field.
	statsDescWins := statsFields[3].Descriptor()
	// stats.DefaultWins holds the default value on creation for the wins field.
	stats.DefaultWins = statsDescWins.Default.(int)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescEliminated is the schema descriptor for eliminated field.
	teamDescEliminated := teamFields[3].Descriptor()
	// team.DefaultEliminated holds the default value on creation for the eliminated field.
	team.DefaultEliminated = teamDescEliminated.Default.(bool)
	// teamDescID is the schema descriptor for id field.
	teamDescID := teamFields[0].Descriptor()
	// team.IDValidator is a validator for the "id" field. It is called by the builders before save.
	team.IDValidator = teamDescID.Validators[0].(func(int) error)
}
