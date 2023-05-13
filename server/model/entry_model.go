package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/entry"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/league"
)

type EntryModel struct {
	client *ent.Client
}

func NewEntryModel(client *ent.Client) *EntryModel {
	return &EntryModel{client: client}
}

func (em *EntryModel) CreateEntry(ownerName string, league *ent.League, players []*ent.Player) (*ent.Entry, error) {
	entry := em.client.Entry.
		Create().
		SetOwnerName(ownerName).
		SetLeague(league).
		AddPlayers(players...)

	e, err := entry.Save(context.Background())
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (em *EntryModel) GetEntryByID(id int) (*ent.Entry, error) {
	return em.client.Entry.
		Get(context.Background(), id)
}

func (em *EntryModel) UpdateEntry(id int, ownerName string, leagueID int, forwardsIDs, defendersIDs, goaliesIDs []int) (*ent.Entry, error) {
	l, err := em.client.League.Get(context.Background(), leagueID)
	if err != nil {
		return nil, errors.New("invalid league ID")
	}

	pm := NewPlayerModel(em.client)

	forwards, err := pm.GetPlayersByID(forwardsIDs)
	defenders, err := pm.GetPlayersByID(defendersIDs)
	goalies, err := pm.GetPlayersByID(goaliesIDs)
	players := append(forwards, defenders...)
	players = append(players, goalies...)
	if err != nil {
		return nil, fmt.Errorf("invalid player IDs: %w", err)
	}

	e, err := em.client.Entry.
		UpdateOneID(id).
		SetOwnerName(ownerName).
		SetLeague(l).
		ClearPlayers().
		AddPlayers(players...).
		Save(context.Background())

	if err != nil {
		return nil, err
	}
	return e, nil
}

func (em *EntryModel) GetEntriesByLeagueID(leagueID int) ([]*ent.Entry, error) {
	entries, err := em.client.Entry.
		Query().
		Where(entry.HasLeagueWith(league.ID(leagueID))).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return entries, nil
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

func (em *EntryModel) GetEntryPlayers(entryID int) ([]*ent.Player, error) {
	entry, err := em.client.Entry.
		Query().
		Where(entry.ID(entryID)).
		WithPlayers().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	players := entry.Edges.Players
	return players, nil
}
