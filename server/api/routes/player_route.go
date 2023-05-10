package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/model"
)

type PlayerAPI struct {
	model *model.PlayerModel
}

func NewPlayerAPI(model *model.PlayerModel) *PlayerAPI {
	return &PlayerAPI{model: model}
}

func (pa *PlayerAPI) RegisterRoutes(app *fiber.App) {
	app.Post("/players", pa.createPlayer)
	app.Get("/players/:id", pa.getPlayerByID)
	app.Put("/players/:id", pa.updatePlayer)
	app.Delete("/players/:id", pa.deletePlayer)
	app.Get("/players", pa.listPlayers)
}

func (pa *PlayerAPI) createPlayer(c *fiber.Ctx) error {
	var player ent.Player
	if err := c.BodyParser(&player); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newPlayer, err := pa.model.CreatePlayer(player.Name, string(player.Position), player.Goals, player.Assists, player.Shutouts, player.Wins, player.ID, player.Edges.Team)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(newPlayer)
}

func (pa *PlayerAPI) getPlayerByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	player, err := pa.model.GetPlayerByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(player)
}

func (pa *PlayerAPI) updatePlayer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var player ent.Player
	if err := c.BodyParser(&player); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedPlayer, err := pa.model.UpdatePlayer(id, player.Name, string(player.Position), player.Goals, player.Assists, player.Shutouts, player.Wins, player.Edges.Team)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedPlayer)
}

func (pa *PlayerAPI) deletePlayer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = pa.model.DeletePlayer(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (pa *PlayerAPI) listPlayers(c *fiber.Ctx) error {
	players, err := pa.model.ListPlayers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(players)
}
