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

func CreateLeague(lm *model.LeagueModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the request body to get the parameters for creating a new league
		type entryData struct {
			OwnerName   string `json:"owner_name"`
			ForwardIDs  []int  `json:"forwards"`
			DefenderIDs []int  `json:"defenders"`
			GoalieIDs   []int  `json:"goalies"`
		}

		var data struct {
			Season       int          `json:"season"`
			IsPublic     bool         `json:"is_public"`
			NumForwards  int          `json:"num_forwards"`
			NumDefenders int          `json:"num_defenders"`
			NumGoalies   int          `json:"num_goalies"`
			Name         string       `json:"name"`
			EditKey      string       `json:"edit_key"`
			Code         string       `json:"code"`
			Entries      []*ent.Entry `json:"entries"`
		}

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Use the LeagueModel to create a new league
		league, err := lm.CreateLeagueWithEntries(
			data.Season,
			data.IsPublic,
			data.NumForwards,
			data.NumDefenders,
			data.NumGoalies,
			data.Name,
			data.EditKey,
			data.Code,
			data.Entries,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(league)
	}
}
