package api

import (
	"strconv"

	"github.com/rj-davidson/stanley-cup-fantasy-hockey/model"

	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"

	"github.com/gofiber/fiber/v2"
)

type LeagueAPI struct {
	model *model.LeagueModel
}

func NewLeagueAPI(model *model.LeagueModel) *LeagueAPI {
	return &LeagueAPI{model: model}
}

func (la *LeagueAPI) RegisterRoutes(app *fiber.App) {
	app.Post("/leagues", la.createLeague)
	app.Get("/leagues/:id", la.getLeagueByID)
	app.Put("/leagues/:id", la.updateLeague)
	app.Delete("/leagues/:id", la.deleteLeague)
	app.Get("/leagues", la.listLeagues)
}

func (la *LeagueAPI) createLeague(c *fiber.Ctx) error {
	var league ent.League
	if err := c.BodyParser(&league); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newLeague, err := la.model.CreateLeague(league.Season, league.Public, league.NumForwards, league.NumDefenders, league.NumGoalies, league.Name, league.EditKey, league.Code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(newLeague)
}

func (la *LeagueAPI) getLeagueByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	league, err := la.model.GetLeagueByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(league)
}

func (la *LeagueAPI) updateLeague(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var league ent.League
	if err := c.BodyParser(&league); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedLeague, err := la.model.UpdateLeague(id, league.Season, league.Public, league.NumForwards, league.NumDefenders, league.NumGoalies, league.EditKey, league.Code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedLeague)
}

func (la *LeagueAPI) deleteLeague(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = la.model.DeleteLeague(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (la *LeagueAPI) listLeagues(c *fiber.Ctx) error {
	leagues, err := la.model.ListLeagues()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(leagues)
}
