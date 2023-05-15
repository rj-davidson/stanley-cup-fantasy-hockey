package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/model"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/utils"
	"io/ioutil"
	"net/http"
	"time"
)

type NHLTeam struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NHLApiResponse struct {
	Teams []NHLTeam `json:"teams"`
}

const nhlApiUrl = "https://statsapi.web.nhl.com/api/v1/teams"

func fetchTeams() ([]NHLTeam, error) {
	httpClient := &http.Client{Timeout: 10 * time.Second}

	resp, err := httpClient.Get(nhlApiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse NHLApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return apiResponse.Teams, nil
}

type TeamController struct {
	teamModel *model.TeamModel
}

func NewTeamController(client *ent.Client) *TeamController {
	return &TeamController{
		teamModel: model.NewTeamModel(client),
	}
}

func (ctrl *TeamController) AddTeams() error {
	nhlTeams, err := fetchTeams()
	if err != nil {
		return err
	}

	loadedTeamIDs, err := ctrl.teamModel.ListTeamIDs()

	for _, nhlTeam := range nhlTeams {
		if utils.Contains(loadedTeamIDs, nhlTeam.ID) {
			fmt.Println("Team already added: %s\n", nhlTeam.Name)
			continue
		} else {
			_, err := ctrl.teamModel.CreateTeam(nhlTeam.Name, fmt.Sprintf("nhl_%d_logo.png", nhlTeam.ID), nhlTeam.ID, true)
			if err != nil {
				return err
			} else {
				fmt.Printf("Added NHL team: %s\n", nhlTeam.Name)
			}
		}
	}

	return nil
}

func (ctrl *TeamController) UpdateEliminatedTeams() {
	// Get all teams who have not been eliminated
	teams, err := ctrl.teamModel.ListPlayoffTeams()
	if err != nil {
		fmt.Printf("Error getting playoff teams: %v", err)
	}
	for _, team := range teams {
		games, err := ctrl.teamModel.ListGamesByTeam(team)
		if err != nil {
			fmt.Printf("Error getting games by team: %v", err)
		}
		// Dict of key: opposingTeam, value: wins
		// If opposingTeam is not in dict, add it with value 0
		oppoWins := make(map[int]int)

		// Search each game
		for _, game := range games {
			awayTeamID := game.QueryAwayTeam().OnlyX(context.Background()).ID
			homeTeamID := game.QueryHomeTeam().OnlyX(context.Background()).ID

			// If team is home team
			if game.QueryAwayTeam().OnlyX(context.Background()) != team && game.AwayScore > game.HomeScore {
				oppoWins[awayTeamID] += 1
			}
			// If team is away team
			if game.QueryHomeTeam().OnlyX(context.Background()) != team && game.HomeScore > game.AwayScore {
				oppoWins[homeTeamID] += 1
			}
		}

		// If any opposing team has 4 wins, team is eliminated
		for _, wins := range oppoWins {
			if wins == 4 {
				_, err := team.Update().SetEliminated(true).Save(context.Background())
				if err != nil {
					fmt.Printf("Error updating team playoff elimination status: %v", err)
				}
			}
		}
	}
}
