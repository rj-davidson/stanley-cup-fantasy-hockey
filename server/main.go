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
		AllowOrigins: "http://" + viper.GetString("IP_HOST") + ":3000",
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
		teamController := controller.NewNHLController(client)
		err := teamController.AddNHLTeams()
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

	api.RegisterLeagueRoutes(app, lm, em, pm)
	api.RegisterPlayerRoutes(app, pm)
}

func updateData(client *ent.Client) {
	// Update playoff games
	updatePlayoffGames(client)

	// Update player points
	updatePlayerPoints(client)

	// Update goalie wins and shutouts
	updateGoalieStats(client)

	fmt.Println("Data update completed")
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

	playerController := controller.NewPlayerController(client)
	for _, team := range teams {
		err := playerController.FetchNHLPlayerPoints(team, context.Background())
		if err != nil {
			fmt.Printf("Error fetching player points: %s\n", err)
		}
	}
}

func updateGoalieStats(client *ent.Client) {
	playerController := controller.NewPlayerController(client)
	err := playerController.UpdateGoalieWinsShutouts()
	if err != nil {
		fmt.Printf("Error updating goalie stats: %s\n", err)
	}
}

func scheduleDataUpdates(client *ent.Client) {
	c := cron.New()

	// Schedule data updates every 10 minutes
	_, err := c.AddFunc("*/10 * * * *", func() {
		fmt.Println("Running data update...")
		updateData(client)
	})
	if err != nil {
		log.Fatalf("failed to schedule data updates: %v", err)
	}

	// Start the cron scheduler
	c.Start()
}
