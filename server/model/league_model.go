package model

import (
	"context"
	"errors"

	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
)

type LeagueModel struct {
	Client *ent.Client
}

func NewLeagueModel(client *ent.Client) *LeagueModel {
	return &LeagueModel{Client: client}
}

func (lm *LeagueModel) CreateLeague(season int, isPublic bool, numForwards, numDefenders, numGoalies int, name string) (*ent.League, error) {
	l, err := lm.Client.League.
		Create().
		SetName(name).
		SetSeason(season).
		SetPublic(isPublic).
		SetNumForwards(numForwards).
		SetNumDefenders(numDefenders).
		SetNumGoalies(numGoalies).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (lm *LeagueModel) CreateLeagueWithEntries(season int, isPublic bool, numForwards, numDefenders, numGoalies int, name string, entries []*ent.Entry) (*ent.League, error) {
	l, err := lm.Client.League.
		Create().
		SetName(name).
		SetSeason(season).
		SetPublic(isPublic).
		SetNumForwards(numForwards).
		SetNumDefenders(numDefenders).
		SetNumGoalies(numGoalies).
		AddEntries(entries...).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (lm *LeagueModel) GetLeagueByID(id int) (*ent.League, error) {
	return lm.Client.League.
		Get(context.Background(), id)
}

func (lm *LeagueModel) UpdateLeague(id int, isPublic bool, numForwards, numDefenders, numGoalies int) (*ent.League, error) {
	l, err := lm.Client.League.
		UpdateOneID(id).
		SetPublic(isPublic).
		SetNumForwards(numForwards).
		SetNumDefenders(numDefenders).
		SetNumGoalies(numGoalies).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (lm *LeagueModel) DeleteLeague(id int) error {
	err := lm.Client.League.
		DeleteOneID(id).
		Exec(context.Background())
	if err != nil {
		return errors.New("failed to delete league")
	}
	return nil
}

func (lm *LeagueModel) ListLeagues() ([]*ent.League, error) {
	return lm.Client.League.
		Query().
		All(context.Background())
}
