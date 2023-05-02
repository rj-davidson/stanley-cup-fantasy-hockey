package model

import (
	"context"
	"errors"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
)

type TeamModel struct {
	client *ent.Client
}

func NewTeamModel(client *ent.Client) *TeamModel {
	return &TeamModel{client: client}
}

func (tm *TeamModel) CreateTeam(name, logoFilepath string, eliminated bool) (*ent.Team, error) {
	return tm.client.Team.
		Create().
		SetName(name).
		SetLogoFilepath(logoFilepath).
		SetEliminated(eliminated).
		Save(context.Background())
}

func (tm *TeamModel) GetTeamByID(id int) (*ent.Team, error) {
	return tm.client.Team.
		Get(context.Background(), id)
}

func (tm *TeamModel) UpdateTeam(id int, name, logoFilepath string, eliminated bool) (*ent.Team, error) {
	t, err := tm.client.Team.
		UpdateOneID(id).
		SetName(name).
		SetLogoFilepath(logoFilepath).
		SetEliminated(eliminated).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (tm *TeamModel) DeleteTeam(id int) error {
	err := tm.client.Team.
		DeleteOneID(id).
		Exec(context.Background())
	if err != nil {
		return errors.New("failed to delete team")
	}
	return nil
}

func (tm *TeamModel) ListTeams() ([]*ent.Team, error) {
	return tm.client.Team.
		Query().
		All(context.Background())
}
