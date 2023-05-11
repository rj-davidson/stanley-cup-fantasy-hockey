package model

import (
	"context"
	"errors"

	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
)

type LeagueModel struct {
	client *ent.Client
}

func NewLeagueModel(client *ent.Client) *LeagueModel {
	return &LeagueModel{client: client}
}

func (lm *LeagueModel) CreateLeague(season int, isPublic bool, numForwards, numDefenders, numGoalies int, name, editKey, code string) (*ent.League, error) {
	l, err := lm.client.League.
		Create().
		SetName(name).
		SetSeason(season).
		SetPublic(isPublic).
		SetNumForwards(numForwards).
		SetNumDefenders(numDefenders).
		SetNumGoalies(numGoalies).
		SetEditKey(editKey).
		SetCode(code).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (lm *LeagueModel) CreateLeagueWithEntries(season int, isPublic bool, numForwards, numDefenders, numGoalies int, name, editKey, code string, entries []*ent.Entry) (*ent.League, error) {
	l, err := lm.client.League.
		Create().
		SetName(name).
		SetSeason(season).
		SetPublic(isPublic).
		SetNumForwards(numForwards).
		SetNumDefenders(numDefenders).
		SetNumGoalies(numGoalies).
		SetEditKey(editKey).
		SetCode(code).
		AddEntries(entries...).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (lm *LeagueModel) GetLeagueByID(id int) (*ent.League, error) {
	return lm.client.League.
		Get(context.Background(), id)
}

func (lm *LeagueModel) UpdateLeague(id, season int, isPublic bool, numForwards, numDefenders, numGoalies int, editKey, code string) (*ent.League, error) {
	l, err := lm.client.League.
		UpdateOneID(id).
		SetPublic(isPublic).
		SetNumForwards(numForwards).
		SetNumDefenders(numDefenders).
		SetNumGoalies(numGoalies).
		SetEditKey(editKey).
		SetCode(code).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (lm *LeagueModel) DeleteLeague(id int) error {
	err := lm.client.League.
		DeleteOneID(id).
		Exec(context.Background())
	if err != nil {
		return errors.New("failed to delete league")
	}
	return nil
}

func (lm *LeagueModel) ListLeagues() ([]*ent.League, error) {
	return lm.client.League.
		Query().
		All(context.Background())
}
