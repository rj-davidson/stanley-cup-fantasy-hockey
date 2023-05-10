package controller_tests

import (
	"context"
	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/controller"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/model"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func TestAddNHLTeams(t *testing.T) {
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
		t.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	nhlController := controller.NewNHLController(client)
	err = nhlController.AddNHLTeams()
	if err != nil {
		t.Fatalf("Error adding NHL teams: %v", err)
	}

	// Create a new TeamModel instance
	teamModel := model.NewTeamModel(client)

	// Query the teams from the database
	teams, err := teamModel.ListTeams()
	if err != nil {
		t.Fatalf("Error querying teams: %v", err)
	}

	// Check if the number of teams in the database matches the expected number of NHL teams
	expectedTeamCount := 32 // Adjust this number according to the expected number of teams
	if len(teams) != expectedTeamCount {
		t.Errorf("Expected %d teams, but got %d", expectedTeamCount, len(teams))
	}

	// Check if a specific team exists in the database
	specificTeamName := "Toronto Maple Leafs"
	foundSpecificTeam := false
	for _, team := range teams {
		if team.Name == specificTeamName {
			foundSpecificTeam = true
			break
		}
	}
	if !foundSpecificTeam {
		t.Errorf("Expected to find team %s, but it was not found", specificTeamName)
	}
}
