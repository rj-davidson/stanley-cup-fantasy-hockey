package model

import (
	"context"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/game"
)

type GameModel struct {
	client *ent.Client
}

func NewGameModel(client *ent.Client) *GameModel {
	return &GameModel{client: client}
}

// CreateGame creates a new game in the database.
func (gm *GameModel) CreateGame(id int, homeWin bool, homeScore, awayScore int, homeTeam, awayTeam *ent.Team, homeGoalie, awayGoalie *ent.Player) (*ent.Game, error) {
	return gm.client.Game.
		Create().
		SetID(id).
		SetHomeWin(homeWin).
		SetHomeScore(homeScore).
		SetAwayScore(awayScore).
		SetHomeTeam(homeTeam).
		SetAwayTeam(awayTeam).
		SetHomeGoalie(homeGoalie).
		SetAwayGoalie(awayGoalie).
		Save(context.Background())
}

// GetGameByID retrieves a game by its ID.
func (gm *GameModel) GetGameByID(id int) (*ent.Game, error) {
	return gm.client.Game.
		Query().
		Where(game.ID(id)).
		WithHomeTeam().
		WithAwayTeam().
		WithHomeGoalie().
		WithAwayGoalie().
		Only(context.Background())
}

// UpdateGame updates a game by its ID.
func (gm *GameModel) UpdateGame(id int, homeWin bool, homeScore, awayScore int) (*ent.Game, error) {
	gameToUpdate, err := gm.GetGameByID(id)
	if err != nil {
		return nil, err
	}

	return gameToUpdate.
		Update().
		SetHomeWin(homeWin).
		SetHomeScore(homeScore).
		SetAwayScore(awayScore).
		Save(context.Background())
}

// DeleteGame deletes a game by its ID.
func (gm *GameModel) DeleteGame(id int) error {
	return gm.client.Game.
		DeleteOneID(id).
		Exec(context.Background())
}
