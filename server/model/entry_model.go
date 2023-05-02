package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
)

type EntryModel struct {
	client *ent.Client
}

func NewEntryModel(client *ent.Client) *EntryModel {
	return &EntryModel{client: client}
}

func (em *EntryModel) CreateEntry(ownerName string, pointTotal int, leagueID int, forwardsIDs, defendersIDs, goaliesIDs []int) (*ent.Entry, error) {
	league, err := em.client.League.Get(context.Background(), leagueID)
	if err != nil {
		return nil, errors.New("invalid league ID")
	}

	pm := NewPlayerModel(em.client)

	forwards, err := pm.GetPlayersByID(forwardsIDs)
	defenders, err := pm.GetPlayersByID(defendersIDs)
	goalies, err := pm.GetPlayersByID(goaliesIDs)
	if err != nil {
		return nil, fmt.Errorf("invalid player IDs: %w", err)
	}

	entry, err := em.client.Entry.
		Create().
		SetOwnerName(ownerName).
		SetPointTotal(pointTotal).
		SetLeague(league).
		AddForwards(forwards...).
		AddDefenders(defenders...).
		AddGoalies(goalies...).
		Save(context.Background())

	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (em *EntryModel) GetEntryByID(id int) (*ent.Entry, error) {
	return em.client.Entry.
		Get(context.Background(), id)
}

func (em *EntryModel) UpdateEntry(id int, ownerName string, pointTotal int, leagueID int, forwardsIDs, defendersIDs, goaliesIDs []int) (*ent.Entry, error) {
	league, err := em.client.League.Get(context.Background(), leagueID)
	if err != nil {
		return nil, errors.New("invalid league ID")
	}

	pm := NewPlayerModel(em.client)

	forwards, err := pm.GetPlayersByID(forwardsIDs)
	defenders, err := pm.GetPlayersByID(defendersIDs)
	goalies, err := pm.GetPlayersByID(goaliesIDs)
	if err != nil {
		return nil, fmt.Errorf("invalid player IDs: %w", err)
	}

	e, err := em.client.Entry.
		UpdateOneID(id).
		SetOwnerName(ownerName).
		SetPointTotal(pointTotal).
		SetLeague(league).
		ClearForwards().
		AddForwards(forwards...).
		ClearDefenders().
		AddDefenders(defenders...).
		ClearGoalies().
		AddGoalies(goalies...).
		Save(context.Background())

	if err != nil {
		return nil, err
	}
	return e, nil
}

func (em *EntryModel) DeleteEntry(id int) error {
	err := em.client.Entry.
		DeleteOneID(id).
		Exec(context.Background())

	if err != nil {
		return errors.New("failed to delete entry")
	}
	return nil
}

func (em *EntryModel) ListEntries() ([]*ent.Entry, error) {
	return em.client.Entry.
		Query().
		All(context.Background())
}
