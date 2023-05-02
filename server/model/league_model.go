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

func (lm *LeagueModel) CreateLeague(season int, isPublic bool, numForwards, numDefenders, numGoalies, pointsForGoal, pointsForAssist, goaliePointsForShutout, goaliePointsForWin int, editKey, code string) (*ent.League, error) {
	l, err := lm.client.League.
		Create().
		SetSeason(season).
		SetPublic(isPublic).
		SetNumForwards(numForwards).
		SetNumDefenders(numDefenders).
		SetNumGoalies(numGoalies).
		SetPointsForGoal(pointsForGoal).
		SetPointsForAssist(pointsForAssist).
		SetGoaliePointsForShutout(goaliePointsForShutout).
		SetGoaliePointsForWin(goaliePointsForWin).
		SetEditKey(editKey).
		SetCode(code).
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

func (lm *LeagueModel) UpdateLeague(id, season int, isPublic bool, numForwards, numDefenders, numGoalies, pointsForGoal, pointsForAssist, goaliePointsForShutout, goaliePointsForWin int, editKey, code string) (*ent.League, error) {
	l, err := lm.client.League.
		UpdateOneID(id).
		SetPublic(isPublic).
		SetNumForwards(numForwards).
		SetNumDefenders(numDefenders).
		SetNumGoalies(numGoalies).
		SetPointsForGoal(pointsForGoal).
		SetPointsForAssist(pointsForAssist).
		SetGoaliePointsForShutout(goaliePointsForShutout).
		SetGoaliePointsForWin(goaliePointsForWin).
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
