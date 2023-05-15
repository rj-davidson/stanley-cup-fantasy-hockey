package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/api/handlers"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/model"
)

func RegisterLeagueRoutes(app *fiber.App, lm *model.LeagueModel, em *model.EntryModel, pm *model.PlayerModel, client *ent.Client) {
	app.Post("/leagues", handlers.CreateLeague(lm, em, pm))
	app.Get("/leagues", handlers.ListLeagues(lm))
	app.Get("/leagues/:id", handlers.GetLeagueByID(client))
	// Add other league routes if needed
}

func RegisterPlayerRoutes(app *fiber.App, pm *model.PlayerModel) {
	app.Get("/players", handlers.ListPlayers(pm))
	app.Get("/players/:id", handlers.GetPlayerByID(pm))
	app.Delete("/players/:id", handlers.DeletePlayer(pm))
	// Add other player routes if needed
}
