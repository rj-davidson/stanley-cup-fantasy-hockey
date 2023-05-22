package main

import (
	"context"
	"flag"
	"fmt"
	api "github.com/rj-davidson/stanley-cup-fantasy-hockey/api/routes"
	"github.com/robfig/cron/v3"
	"log"

	"entgo.io/ent/dialect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/controller"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent/migrate"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/model"
	"github.com/spf13/viper"
)

var initData bool

func init() {
	flag.BoolVar(&initData, "init-data", false, "Initialize data on startup")
	flag.Parse()
}

func main() {
	// Set up Viper configuration
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	viper.AutomaticEnv()

	// Build Connection String
	connectionString := "postgres://" + viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASS") + "@" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + "/" + viper.GetString("DB_NAME") + "?sslmode=disable"

	// Create a new Ent client instance
	client, err := ent.Open(dialect.Postgres, connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run migrations on startup
	err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Web App
	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://" + viper.GetString("IP_HOST") +
			":3000, http://localhost:3000, https://" + viper.GetString("DOMAIN") + ", http://" +
			viper.GetString("IP HOST") + ":80, http://" + viper.GetString("IP_HOST"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Initialize data
	if initData {
		initializeData(client) // Use the exported function
	}

	// Set up routes
	initRoutes(app, client)

	// Schedule data updates every 10 minutes
	go scheduleDataUpdates(client)

	fmt.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}

func initializeData(client *ent.Client) {
	// Add teams
	{
		teamController := controller.NewTeamController(client)
		err := teamController.AddTeams()
		if err != nil {
			fmt.Println(err)
		}
	}

	// Add players
	{
		teamModel := model.NewTeamModel(client)
		nhlTeams, _ := teamModel.ListTeams()
		playerController := controller.NewPlayerController(client)
		err := playerController.AddNHLPlayers(nhlTeams)
		if err != nil {
			fmt.Printf("Error adding players: %s", err)
		}
	}

	updateData(client)
}

func initRoutes(app *fiber.App, client *ent.Client) {
	lm := model.NewLeagueModel(client)
	pm := model.NewPlayerModel(client)
	em := model.NewEntryModel(client)

	api.RegisterLeagueRoutes(app, lm, em, pm, client)
	api.RegisterPlayerRoutes(app, pm)
}

func updateData(client *ent.Client) {
	// Update playoff games
	updatePlayoffGames(client)
	fmt.Println("[Update] playoff games completed")

	// Update player points
	updatePlayerPoints(client)
	fmt.Println("[Update] player points completed")

	// Update eliminated teams
	updateTeamEliminated(client)
	fmt.Println("[Update] eliminated teams completed")

	fmt.Println("-- DATA UPDATE COMPLETE --")
}

func updatePlayoffGames(client *ent.Client) {
	gameController := controller.NewGameController(client)
	err := gameController.FetchPostSeasonGames()
	if err != nil {
		fmt.Printf("Error fetching playoff games: %s", err)
	}
}

func updatePlayerPoints(client *ent.Client) {
	// Fetch all teams
	teamModel := model.NewTeamModel(client)
	teams, err := teamModel.ListPlayoffTeams()
	if err != nil {
		fmt.Printf("Error fetching teams: %s\n", err)
		return
	}

	// Get all players on playoff teams
	playerModel := model.NewPlayerModel(client)
	players, err := playerModel.ListPlayersOnTeams(teams)
	if err != nil {
		fmt.Printf("Error fetching players: %s\n", err)
		return
	}
	updatePlayerStats(client, players)
}

func updatePlayerStats(client *ent.Client, players []*ent.Player) {
	playerController := controller.NewPlayerController(client)
	err := playerController.UpdateStats(players)
	if err != nil {
		fmt.Printf("Error updating player stats: %s\n", err)
	}
}

func updateTeamEliminated(client *ent.Client) {
	teamController := controller.NewTeamController(client)
	teamController.UpdateEliminatedTeams()
}

func scheduleDataUpdates(client *ent.Client) {
	c := cron.New()

	// Schedule data updates every 20 minutes
	_, err := c.AddFunc("*/20 * * * *", func() {
		fmt.Println("Running data update...")
		updateData(client)
	})
	if err != nil {
		log.Fatalf("failed to schedule data updates: %v", err)
	}

	// Start the cron scheduler
	c.Start()
}
