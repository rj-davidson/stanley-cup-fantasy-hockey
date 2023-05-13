package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/model"
)

func ListPlayers(pm *model.PlayerModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		players, err := pm.ListPlayers()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(players)
	}
}

func GetPlayerByID(pm *model.PlayerModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		player, err := pm.GetPlayerByID(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(player)
	}
}

func DeletePlayer(pm *model.PlayerModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		err = pm.DeletePlayer(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
