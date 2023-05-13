package handlers

import (
	"fmt"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
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

func GetLeagueByID(lm *model.LeagueModel, em *model.EntryModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		league, err := lm.GetLeagueByID(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		// Create a response structure to hold the league data
		response := LeagueStruct{
			ID:           league.ID,
			Name:         league.Name,
			Season:       league.Season,
			Public:       league.Public,
			NumForwards:  league.NumForwards,
			NumDefenders: league.NumDefenders,
			NumGoalies:   league.NumGoalies,
			Entries:      []EntryStruct{},
		}

		// Retrieve the entries for the league
		entries, err := em.GetEntriesByLeagueID(league.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Iterate over the entries and populate the response structure
		for _, entry := range entries {
			entryData := EntryStruct{
				ID:        entry.ID,
				OwnerName: entry.OwnerName,
				Players:   []PlayerStruct{},
			}

			// Fetch all rows from the "entry_players" join table with entry.ID equals to the entry ID
			entryPlayers, err := em.GetEntryPlayers(entry.ID)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}

			for _, player := range entryPlayers {
				playerData := PlayerStruct{
					ID:       player.ID,
					Name:     player.Name,
					Position: player.Position.String(),
					Goals:    player.Goals,
					Assists:  player.Assists,
					Shutouts: player.Shutouts,
					Wins:     player.Wins,
				}
				entryData.Players = append(entryData.Players, playerData)
			}

			response.Entries = append(response.Entries, entryData)
		}

		return c.JSON(response)
	}
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
	Team     string `json:"team"`
}

func validatePlayers(pm *model.PlayerModel, playerData []PlayerStruct) ([]*ent.Player, error) {
	playerIDs := make([]int, len(playerData))
	for i, player := range playerData {
		playerIDs[i] = player.ID
	}

	return pm.GetPlayersByID(playerIDs)
}

func CreateLeague(lm *model.LeagueModel, em *model.EntryModel, pm *model.PlayerModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the request body
		var leagueData LeagueStruct
		err := c.BodyParser(&leagueData)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Create a new league entity
		newLeague, err := lm.CreateLeague(
			leagueData.Season,
			leagueData.Public,
			leagueData.NumForwards,
			leagueData.NumDefenders,
			leagueData.NumGoalies,
			leagueData.Name,
		)
		fmt.Printf("Created league: %+v\n", newLeague)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create league"})
		}

		// Iterate over the entries and add them to the league
		for _, entryData := range leagueData.Entries {
			// Validate the players by their IDs
			players, err := validatePlayers(pm, entryData.Players)

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
