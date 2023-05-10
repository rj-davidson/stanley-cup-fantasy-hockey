package model

import (
	"context"
	"errors"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/player"
)

type PlayerModel struct {
	client *ent.Client
}

func NewPlayerModel(client *ent.Client) *PlayerModel {
	return &PlayerModel{client: client}
}

func (pm *PlayerModel) CreatePlayer(name, position string, goals, assists, shutouts, wins, id int, team *ent.Team) (*ent.Player, error) {
	return pm.client.Player.
		Create().
		SetID(id).
		SetName(name).
		SetPosition(player.Position(position)).
		SetGoals(goals).
		SetAssists(assists).
		SetShutouts(shutouts).
		SetWins(wins).
		SetTeam(team).
		Save(context.Background())
}

func (pm *PlayerModel) GetPlayerByID(id int) (*ent.Player, error) {
	return pm.client.Player.
		Get(context.Background(), id)
}

func (pm *PlayerModel) GetPlayersByID(ids []int) ([]*ent.Player, error) {
	return pm.client.Player.
		Query().
		Where(player.IDIn(ids...)).
		All(context.Background())
}

func (pm *PlayerModel) UpdatePlayer(id int, name, position string, goals, assists, shutouts, wins int, team *ent.Team) (*ent.Player, error) {
	p, err := pm.client.Player.
		UpdateOneID(id).
		SetName(name).
		SetPosition(player.Position(position)).
		SetGoals(goals).
		SetAssists(assists).
		SetShutouts(shutouts).
		SetWins(wins).
		SetTeam(team).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (pm *PlayerModel) DeletePlayer(id int) error {
	err := pm.client.Player.
		DeleteOneID(id).
		Exec(context.Background())
	if err != nil {
		return errors.New("failed to delete player")
	}
	return nil
}

func (pm *PlayerModel) ListPlayers() ([]*ent.Player, error) {
	return pm.client.Player.
		Query().
		All(context.Background())
}

func (pm *PlayerModel) ListPlayerIDs() ([]int, error) {
	players, err := pm.client.Player.
		Query().
		All(context.Background())

	if err != nil {
		return nil, err
	}

	playerIDs := make([]int, len(players))
	for i, player := range players {
		playerIDs[i] = player.ID
	}

	return playerIDs, nil
}
