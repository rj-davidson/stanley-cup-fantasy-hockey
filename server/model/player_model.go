package model

import (
	"context"
	"errors"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/entry"
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

// Function for updating player stats
func (pm *PlayerModel) UpdatePlayerPoints(player *ent.Player, goals, assists int) (*ent.Player, error) {
	p, err := pm.client.Player.
		UpdateOne(player).
		SetGoals(goals).
		SetAssists(assists).
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

// UpdateGoalieWins updates the wins for a goalie
func (pm *PlayerModel) UpdateGoalieWins(goalie *ent.Player, wins int) (*ent.Player, error) {
	p, err := pm.client.Player.
		UpdateOne(goalie).
		SetWins(wins).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return p, nil
}

// UpdateGoalieShutouts updates the shutouts for a goalie
func (pm *PlayerModel) UpdateGoalieShutouts(goalie *ent.Player, shutouts int) (*ent.Player, error) {
	p, err := pm.client.Player.
		UpdateOne(goalie).
		SetShutouts(shutouts).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return p, nil
}

// GetPlayersByEntryIDs returns set of unique players for a list of entries
func (pm *PlayerModel) GetPlayersByEntries(entries []*ent.Entry) ([]*ent.Player, error) {
	var IDs []int
	for _, e := range entries {
		IDs = append(IDs, e.ID)
	}
	players, err := pm.client.Player.
		Query().
		Where(player.HasEntriesWith(entry.IDIn(IDs...))).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return players, nil
}
