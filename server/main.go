package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/rj-davidson/stanley-cup-fantasy-hockey/api/handlers"
	api "github.com/rj-davidson/stanley-cup-fantasy-hockey/api/routes"

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
	connectionString := "postgres://" + viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASS") + "@" + viper.GetString("DB_HOST_LOCAL") + ":" + viper.GetString("DB_PORT") + "/" + viper.GetString("DB_NAME") + "?sslmode=disable"

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
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Initialize data
	if initData {
		initializeData(client) // Use the exported function
	}

	// Set up routes
	initializeRoutes(app, client)

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
}

func initializeRoutes(app *fiber.App, client *ent.Client) {
	lm := model.NewLeagueModel(client)
	la := api.NewLeagueAPI(lm)

	la.RegisterRoutes(app)
	app.Get("/leagues", handlers.ListLeagues(lm))
	app.Get("/leagues/:id", handlers.GetLeagueByID(lm))
	app.Post("/leagues", handlers.CreateLeague(lm))

	pm := model.NewPlayerModel(client)
	pa := api.NewPlayerAPI(pm)

	pa.RegisterRoutes(app)
	app.Get("/players", handlers.ListPlayers(pm))
	app.Get("/players/:id", handlers.GetPlayerByID(pm))
}
