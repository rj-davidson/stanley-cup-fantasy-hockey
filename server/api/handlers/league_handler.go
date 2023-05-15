package handlers

import (
	"context"
	"fmt"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/league"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/model"
)

func ListLeagues(lm *model.LeagueModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		leagues, err := lm.ListLeagues()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(leagues)
	}
}

func GetLeagueByID(client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		l, err := client.League.
			Query().
			Where(league.ID(id)).
			WithEntries(func(query *ent.EntryQuery) {
				query.WithPlayers(func(query *ent.PlayerQuery) {
					query.WithTeam()
				})
			}).
			Only(context.Background())

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		// Create a set of type team struct (no duplicates)
		teamSet := make(map[int]TeamStruct)
		playerSet := make(map[int]PlayerStruct)
		var response LeagueBundle
		response.League = LeagueStruct{
			ID:           l.ID,
			Name:         l.Name,
			Season:       l.Season,
			Public:       l.Public,
			NumForwards:  l.NumForwards,
			NumDefenders: l.NumDefenders,
			NumGoalies:   l.NumGoalies,
		}
		for _, e := range l.Edges.Entries {
			var pIDs []int
			for _, p := range e.Edges.Players {
				// Add Teams to the team set
				teamSet[p.Edges.Team.ID] = TeamStruct{
					ID:         p.Edges.Team.ID,
					Name:       p.Edges.Team.Name,
					Eliminated: p.Edges.Team.Eliminated,
				}
				playerSet[p.ID] = PlayerStruct{
					ID:       p.ID,
					Name:     p.Name,
					Position: p.Position.String(),
					Goals:    p.Goals,
					Assists:  p.Assists,
					Shutouts: p.Shutouts,
					Wins:     p.Wins,
					TeamID:   p.Edges.Team.ID,
				}
				pIDs = append(pIDs, p.ID)
			}
			response.Entries = append(response.Entries, EntryResponse{
				ID:        e.ID,
				OwnerName: e.OwnerName,
				PlayerIDs: pIDs,
			})
		}
		response.Players = make([]PlayerStruct, 0, len(playerSet))
		for _, p := range playerSet {
			response.Players = append(response.Players, p)
		}
		response.Teams = make([]TeamStruct, 0, len(teamSet))
		for _, t := range teamSet {
			response.Teams = append(response.Teams, t)
		}
		return c.JSON(response)
	}
}

type EntryResponse struct {
	ID        int    `json:"id,omitempty"`
	OwnerName string `json:"owner_name"`
	PlayerIDs []int  `json:"playerIDs,omitempty"`
}

type LeagueBundle struct {
	League  LeagueStruct    `json:"league"`
	Players []PlayerStruct  `json:"players"`
	Entries []EntryResponse `json:"entries"`
	Teams   []TeamStruct    `json:"teams"`
}

type LeagueStruct struct {
	ID           int           `json:"id,omitempty"`
	Name         string        `json:"name"`
	Season       int           `json:"season"`
	Public       bool          `json:"public"`
	NumForwards  int           `json:"num_forwards"`
	NumDefenders int           `json:"num_defenders"`
	NumGoalies   int           `json:"num_goalies"`
	Entries      []EntryStruct `json:"entries"`
}

type EntryStruct struct {
	ID        int            `json:"id,omitempty"`
	OwnerName string         `json:"owner_name"`
	Players   []PlayerStruct `json:"players,omitempty"`
}

type PlayerStruct struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Goals    int    `json:"goals"`
	Assists  int    `json:"assists"`
	Shutouts int    `json:"shutouts"`
	Wins     int    `json:"wins"`
	TeamID   int    `json:"team_id"`
}

type TeamStruct struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Eliminated bool   `json:"eliminated"`
}

func CreateLeague(lm *model.LeagueModel, em *model.EntryModel, pm *model.PlayerModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the request body
		var leagueBundle LeagueBundle
		err := c.BodyParser(&leagueBundle)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Create a new league entity
		newLeague, err := lm.CreateLeague(
			leagueBundle.League.Season,
			leagueBundle.League.Public,
			leagueBundle.League.NumForwards,
			leagueBundle.League.NumDefenders,
			leagueBundle.League.NumGoalies,
			leagueBundle.League.Name,
		)
		fmt.Printf("Created league: %+v\n", newLeague)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create league"})
		}

		// Iterate over the entries and add them to the league
		for _, entryData := range leagueBundle.Entries {
			// Validate the players by their IDs
			players, err := pm.GetPlayersByID(entryData.PlayerIDs)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch forwards"})
			}

			// Create a new entry entity
			_, err = em.CreateEntry(
				entryData.OwnerName,
				newLeague,
				players,
			)
			fmt.Printf("Created entry: %+v\n", newLeague)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create entry"})
			}
		}

		// Return a success response
		return c.JSON(fiber.Map{
			"message": "League created successfully",
			"league":  newLeague,
		})
	}
}
