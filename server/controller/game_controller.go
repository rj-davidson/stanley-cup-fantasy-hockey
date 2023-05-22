package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/model"
	"io/ioutil"
	"net/http"
	"time"
)

type GameController struct {
	gameModel   *model.GameModel
	playerModel *model.PlayerModel
	teamModel   *model.TeamModel
	statsModel  *model.HockeyStatsModel
}

func NewGameController(client *ent.Client) *GameController {
	return &GameController{
		gameModel:   model.NewGameModel(client),
		playerModel: model.NewPlayerModel(client),
		teamModel:   model.NewTeamModel(client),
		statsModel:  model.NewHockeyStatsModel(client),
	}
}

func (ctrl *GameController) FetchPostSeasonGames() error {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	const nhlApiScheduleURLFormat = "https://statsapi.web.nhl.com/api/v1/schedule?startDate=2023-04-16&endDate=2023-07-10"

	resp, err := httpClient.Get(nhlApiScheduleURLFormat)
	if err != nil {
		return fmt.Errorf("error fetching post season schedule: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	var scheduleResult map[string]interface{}
	if err := json.Unmarshal(body, &scheduleResult); err != nil {
		return fmt.Errorf("error unmarshaling schedule result: %w", err)
	}

	for _, date := range scheduleResult["dates"].([]interface{}) {
		for _, game := range date.(map[string]interface{})["games"].([]interface{}) {
			gameMap := game.(map[string]interface{})
			gameID := int(gameMap["gamePk"].(float64))

			existingGame, _ := ctrl.gameModel.GetGameByID(gameID)

			if existingGame == nil {
				homeTeamID := int(gameMap["teams"].(map[string]interface{})["home"].(map[string]interface{})["team"].(map[string]interface{})["id"].(float64))
				awayTeamID := int(gameMap["teams"].(map[string]interface{})["away"].(map[string]interface{})["team"].(map[string]interface{})["id"].(float64))

				if gameMap["status"].(map[string]interface{})["detailedState"].(string) != "Final" {
					continue
				}

				homeScore := int(gameMap["teams"].(map[string]interface{})["home"].(map[string]interface{})["score"].(float64))
				awayScore := int(gameMap["teams"].(map[string]interface{})["away"].(map[string]interface{})["score"].(float64))

				homeTeam, err := ctrl.teamModel.GetTeamByID(homeTeamID)
				if err != nil {
					return fmt.Errorf("error getting home team: %w", err)
				}
				if err := ctrl.teamModel.SetEliminatedStatus(homeTeam, false); err != nil {
					return fmt.Errorf("error setting home team as playoff competitor: %w", err)
				}

				awayTeam, err := ctrl.teamModel.GetTeamByID(awayTeamID)
				if err != nil {
					return fmt.Errorf("error getting away team: %w", err)
				}
				if err := ctrl.teamModel.SetEliminatedStatus(awayTeam, false); err != nil {
					return fmt.Errorf("error setting away team as playoff competitor: %w", err)
				}

				// Create Game Entity
				// Create game entity
				g, err := ctrl.gameModel.CreateGame(gameID, homeScore, awayScore, homeTeam, awayTeam)
				if err != nil {
					fmt.Printf("Error creating game entity: %v\n", err)
					return err
				}
				fmt.Printf("Fetched game %d\n", gameID)

				boxscoreData, err := ctrl.fetchBoxscore(gameID)
				if err != nil {
					return fmt.Errorf("error fetching boxscore data: %w", err)
				}
				err = ctrl.setGameStats(boxscoreData, g)
				if err != nil {
					return fmt.Errorf("error upserting game stats: %w", err)
				}

				if err != nil {
					return fmt.Errorf("error fetching game data %d: %w", gameID, err)
				}
			} else {
				fmt.Printf("Game %d already exists in the database\n", gameID)
			}
		}
	}
	fmt.Println("----- Successfully fetched post season games")
	return nil
}

func (ctrl *GameController) fetchBoxscore(gameID int) (map[string]interface{}, error) {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	const nhlApiBoxScoreURLFormat = "https://statsapi.web.nhl.com/api/v1/game/%d/boxscore"
	boxScoreURL := fmt.Sprintf(nhlApiBoxScoreURLFormat, gameID)
	resp, err := httpClient.Get(boxScoreURL)
	if err != nil {
		fmt.Printf("Error fetching box score for game %d: %s\n", gameID, err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body for game %d: %s\n", gameID, err.Error())
		return nil, err
	}

	var boxscoreResult map[string]interface{}
	json.Unmarshal([]byte(body), &boxscoreResult)

	return boxscoreResult, nil
}

func (ctrl *GameController) setGameStats(boxscoreData map[string]interface{}, game *ent.Game) error {
	// Get players for home and away teams
	homePlayers := boxscoreData["teams"].(map[string]interface{})["home"].(map[string]interface{})["players"].(map[string]interface{})
	awayPlayers := boxscoreData["teams"].(map[string]interface{})["away"].(map[string]interface{})["players"].(map[string]interface{})

	// Check if home team won the game
	homeWon := false
	shutout := false
	if game.HomeScore > game.AwayScore {
		homeWon = true
		if game.AwayScore == 0 {
			shutout = true
		}
	} else {
		if game.HomeScore == 0 {
			shutout = true
		}
	}

	// Loop through home players
	for _, player := range homePlayers {
		playerMap := player.(map[string]interface{})
		playerID := int(playerMap["person"].(map[string]interface{})["id"].(float64))
		playerPosition := playerMap["position"].(map[string]interface{})["abbreviation"].(string)

		var playerStats map[string]interface{}
		if playerPosition == "G" {
			playerStats = playerMap["stats"].(map[string]interface{})["goalieStats"].(map[string]interface{})
			// If time on ice is 0, goalie did not play
			if playerStats["timeOnIce"].(string) == "0:00" {
				continue
			}
		} else {
			playerStats = playerMap["stats"].(map[string]interface{})["skaterStats"].(map[string]interface{})
		}

		// Create GameStat entity for home team
		_, err := ctrl.statsModel.CreateGameStat(game, true, shutout, homeWon, playerID, int(playerStats["goals"].(float64)), int(playerStats["assists"].(float64)), context.Background())
		if err != nil {
			return fmt.Errorf("error creating game stat entity: %w", err)
		}
	}

	// Loop through away players
	for _, player := range awayPlayers {
		playerMap := player.(map[string]interface{})
		playerID := int(playerMap["person"].(map[string]interface{})["id"].(float64))
		playerPosition := playerMap["position"].(map[string]interface{})["abbreviation"].(string)

		var playerStats map[string]interface{}
		if playerPosition == "G" {
			playerStats = playerMap["stats"].(map[string]interface{})["goalieStats"].(map[string]interface{})
			// If time on ice is 0, goalie did not play
			if playerStats["timeOnIce"].(string) == "0:00" {
				continue
			}
		} else {
			playerStats = playerMap["stats"].(map[string]interface{})["skaterStats"].(map[string]interface{})
		}

		// Create GameStat entity for away team
		_, err := ctrl.statsModel.CreateGameStat(game, false, shutout, !homeWon, playerID, int(playerStats["goals"].(float64)), int(playerStats["assists"].(float64)), context.Background())
		if err != nil {
			return fmt.Errorf("error creating game stat entity: %w", err)
		}
	}

	return nil
}
