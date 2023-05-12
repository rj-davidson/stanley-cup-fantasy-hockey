package handlers

import (
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

func GetLeagueByID(lm *model.LeagueModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		league, err := lm.GetLeagueByID(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(league)
	}
}

type EntryPlayerData struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Position string   `json:"position"`
	Edges    struct{} `json:"edges"`
}

type EntryData struct {
	OwnerName string            `json:"owner_name"`
	Forwards  []EntryPlayerData `json:"forwards"`
	Defenders []EntryPlayerData `json:"defenders"`
	Goalies   []EntryPlayerData `json:"goalies"`
}

type LeagueRequest struct {
	Name         string      `json:"name"`
	Season       int         `json:"season"`
	Public       bool        `json:"public"`
	NumForwards  int         `json:"num_forwards"`
	NumDefenders int         `json:"num_defenders"`
	NumGoalies   int         `json:"num_goalies"`
	Entries      []EntryData `json:"entries"`
}

func validatePlayers(pm *model.PlayerModel, playerData []EntryPlayerData) ([]*ent.Player, error) {
	playerIDs := make([]int, len(playerData))
	for i, player := range playerData {
		playerIDs[i] = player.ID
	}

	return pm.GetPlayersByID(playerIDs)
}

func CreateLeague(lm *model.LeagueModel, em *model.EntryModel, pm *model.PlayerModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the request body
		var leagueData LeagueRequest
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
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create league"})
		}

		// Iterate over the entries and add them to the league
		for _, entryData := range leagueData.Entries {
			// Validate the players by their IDs
			forwards, err := validatePlayers(pm, entryData.Forwards)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch forwards"})
			}

			defenders, err := validatePlayers(pm, entryData.Defenders)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch defenders"})
			}

			goalies, err := validatePlayers(pm, entryData.Goalies)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch goalies"})
			}

			// Create a new entry entity
			_, err = em.CreateEntry(
				entryData.OwnerName,
				newLeague,
				forwards,
				defenders,
				goalies,
			)
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
