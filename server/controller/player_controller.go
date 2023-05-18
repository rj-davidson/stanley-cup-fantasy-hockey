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

type Person struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
}

type Position struct {
	Type string `json:"type"`
}

type NHLPlayer struct {
	Person   Person   `json:"person"`
	Position Position `json:"position"`
}

type NHLRoster struct {
	Roster []NHLPlayer `json:"roster"`
}

const nhlApiRosterUrlFormat = "https://statsapi.web.nhl.com/api/v1/teams/%d/roster?expand=person.stats&stats=statsSingleSeasonPlayoffs&season=%s"
const season = "20222023"

func fetchNHLPlayers(teamID int) ([]NHLPlayer, error) {
	httpClient := &http.Client{Timeout: 10 * time.Second}

	rosterUrl := fmt.Sprintf(nhlApiRosterUrlFormat, teamID, season)
	resp, err := httpClient.Get(rosterUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse NHLRoster
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	players := make([]NHLPlayer, len(apiResponse.Roster))
	for i, person := range apiResponse.Roster {
		players[i] = person
	}

	return players, nil
}

type PlayerController struct {
	playerModel     *model.PlayerModel
	gameModel       *model.GameModel
	HockeyStatModel *model.HockeyStatsModel
}

func NewPlayerController(client *ent.Client) *PlayerController {
	return &PlayerController{
		playerModel: model.NewPlayerModel(client),
		gameModel:   model.NewGameModel(client),
	}
}

func (ctrl *PlayerController) AddNHLPlayers(teams []*ent.Team) error {
	loadedPlayerIDs, _ := ctrl.playerModel.ListPlayerIDs()
	for _, team := range teams {
		nhlPlayers, err := fetchNHLPlayers(team.ID)
		if err != nil {
			return err
		}

		for _, nhlPlayer := range nhlPlayers {
			if utils.Contains(loadedPlayerIDs, nhlPlayer.Person.ID) {
				fmt.Printf("Player %s already loaded\n", nhlPlayer.Person.FullName)
				continue
			} else {
				_, err := ctrl.playerModel.CreatePlayer(
					nhlPlayer.Person.FullName,
					nhlPlayer.Position.Type,
					nhlPlayer.Person.ID,
					team,
				)
				if err != nil {
					fmt.Printf("Error creating player %s: %s\n", nhlPlayer.Person.FullName, err.Error())
				} else {
					fmt.Printf("Created player %s\n", nhlPlayer.Person.FullName)
				}
			}
		}
	}

	return nil
}

func (ctrl *PlayerController) UpdateStats() error {
	players, err := ctrl.playerModel.ListPlayers()
	if err != nil {
		fmt.Printf("Error listing players: %s\n", err.Error())
		return err
	}
	for _, player := range players {
		fmt.Printf("Updating stats for player %s\n", player.Name)
		err = ctrl.HockeyStatModel.UpdatePlayerStats(player, context.Background())
		if err != nil {
			fmt.Printf("Error updating stats for player %s: %s\n", player.Name, err.Error())
			return err
		}
	}
	return err
}
