package model

import (
	"context"
	"fmt"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/game"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/gamestats"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/player"
)

type HockeyStatsModel struct {
	client *ent.Client
}

func NewHockeyStatsModel(client *ent.Client) *HockeyStatsModel {
	return &HockeyStatsModel{client: client}
}

func (hsm *HockeyStatsModel) CreateGameStat(atHome, shutout, win bool, gameID, playerID, goals, assists int, ctx context.Context) (*ent.GameStats, error) {
	p, err := hsm.client.Player.Query().Where(player.IDEQ(playerID)).Only(ctx)
	if err != nil {
		fmt.Printf("error getting player %d: %v\n", playerID, err)
		return nil, err
	}

	g, err := hsm.client.Game.Query().Where(game.IDEQ(gameID)).Only(ctx)
	if err != nil {
		fmt.Printf("error getting game %d: %v\n", gameID, err)
		return nil, err
	}

	gameStat, err := hsm.client.GameStats.
		Create().
		SetPlayer(p).
		SetGame(g).
		SetHomeGame(atHome).
		SetGoals(goals).
		SetAssists(assists).
		SetShutout(shutout).
		SetWin(win).
		Save(ctx)
	return gameStat, err
}

// UpdateGameStat updates the game stats for a player in a game
func (hsm *HockeyStatsModel) UpdateGameStat(gs *ent.GameStats, atHome, shutout, win bool, goals, assists int, ctx context.Context) (*ent.GameStats, error) {
	return hsm.client.GameStats.UpdateOneID(gs.ID).
		SetHomeGame(atHome).
		SetGoals(goals).
		SetAssists(assists).
		SetShutout(shutout).
		SetWin(win).
		Save(ctx)
}

// GetGameStatsByPlayer returns all game stats for a player
func (hsm *HockeyStatsModel) GetGameStatsByPlayer(p *ent.Player, ctx context.Context) ([]*ent.GameStats, error) {
	return hsm.client.Player.QueryGameStats(p).All(ctx)
}

// GetGameStatsByGame returns all game stats for a game
func (hsm *HockeyStatsModel) GetGameStatsByGame(game *ent.Game, ctx context.Context) ([]*ent.GameStats, error) {
	return hsm.client.Game.QueryGameStats(game).All(ctx)
}

// GetGameStatsByPlayerAndGame returns the game stats for a player in a game
func (hsm *HockeyStatsModel) GetGameStatsByPlayerAndGame(p *ent.Player, g *ent.Game, ctx context.Context) (*ent.GameStats, error) {
	return hsm.client.GameStats.Query().Where(
		gamestats.HasPlayerWith(player.IDEQ(p.ID)),
		gamestats.HasGameWith(game.IDEQ(g.ID)),
	).Only(ctx)
}

// UpdatePlayerStats creates or updates a player's stats
func (hsm *HockeyStatsModel) UpdatePlayerStats(player *ent.Player, ctx context.Context) error {
	playerGameStats, err := hsm.GetGameStatsByPlayer(player, ctx)
	if err != nil {
		return err
	}

	goals, assists, wins, shutouts := 0, 0, 0, 0
	for _, gameStat := range playerGameStats {
		goals += gameStat.Goals
		assists += gameStat.Assists
		if gameStat.Win {
			wins++
		}
		if gameStat.Shutout {
			shutouts++
		}
	}

	playerStat, err := hsm.client.Player.QueryStats(player).Only(ctx)
	if err != nil {
		_, err = hsm.client.Stats.Create().
			SetPlayer(player).
			SetGoals(goals).
			SetAssists(assists).
			SetWins(wins).
			SetShutouts(shutouts).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err = hsm.client.Stats.UpdateOneID(playerStat.ID).
			SetGoals(goals).
			SetAssists(assists).
			SetWins(wins).
			SetShutouts(shutouts).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return err
}
