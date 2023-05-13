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

func (ctrl *PlayerController) FetchNHLPlayerPoints(team *ent.Team, ctx context.Context) error {
	type GameType struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Postseason  bool   `json:"postseason"`
	}

	type Type struct {
		DisplayName string   `json:"displayName"`
		GameType    GameType `json:"gameType"`
	}

	type PlayerStat struct {
		Goals   int `json:"goals"`
		Assists int `json:"assists"`
	}

	type Split struct {
		Season string     `json:"season"`
		Stat   PlayerStat `json:"stat"`
	}

	type Stat struct {
		Type   Type    `json:"type"`
		Splits []Split `json:"splits"`
	}

	type Response struct {
		Copyright string `json:"copyright"`
		Stats     []Stat `json:"stats"`
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}
	const nhlApiPlayerStatsURLFormat = "https://statsapi.web.nhl.com/api/v1/people/%d/stats?stats=statsSingleSeasonPlayoffs&season=%s"
	const currentSeason = "20222023"

	for _, player := range team.QueryPlayers().AllX(ctx) {
		statURL := fmt.Sprintf(nhlApiPlayerStatsURLFormat, player.ID, currentSeason)
		resp, err := httpClient.Get(statURL)
		if err != nil {
			fmt.Printf("Error fetching player stats for %s: %s\n", player.Name, err.Error())
			continue
		}
		defer resp.Body.Close()

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			fmt.Printf("Error decoding response for %s: %s\n", player.Name, err.Error())
			continue
		}

		// Assuming there's only one "stats" element in the response
		if len(response.Stats) > 0 {
			stats := response.Stats[0]
			if len(stats.Splits) > 0 {
				playerStat := stats.Splits[0].Stat
				goals := playerStat.Goals
				assists := playerStat.Assists

				ctrl.playerModel.UpdatePlayerPoints(player, goals, assists)
				fmt.Printf("Updated player %s with %d goals and %d assists\n", player.Name, goals, assists)
			}
		}
	}

	return nil
}

type PlayerController struct {
	playerModel *model.PlayerModel
	gameModel   *model.GameModel
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
					0,
					0,
					0,
					0,
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

func (ctrl *PlayerController) UpdateGoalieWinsShutouts() error {
	goalieWins := make(map[int]int)     // Dictionary of key: playerID, value: wins
	goalieShutouts := make(map[int]int) // Dictionary of key: playerID, value: shutouts

	games, err := ctrl.gameModel.ListGames()
	if err != nil {
		return err
	}

	for _, game := range games {
		homeGoalie := game.QueryHomeGoalie().OnlyX(context.Background())
		awayGoalie := game.QueryAwayGoalie().OnlyX(context.Background())

		// Update wins for home goalie
		if homeGoalie != nil {
			if _, ok := goalieWins[homeGoalie.ID]; !ok {
				goalieWins[homeGoalie.ID] = 0
			}
			if game.HomeScore > game.AwayScore {
				goalieWins[homeGoalie.ID]++
			}
		}

		// Update wins for away goalie
		if awayGoalie != nil {
			if _, ok := goalieWins[awayGoalie.ID]; !ok {
				goalieWins[awayGoalie.ID] = 0
			}
			if game.AwayScore > game.HomeScore {
				goalieWins[awayGoalie.ID]++
			}
		}

		// Update shutouts for home goalie
		if homeGoalie != nil {
			if _, ok := goalieShutouts[homeGoalie.ID]; !ok {
				goalieShutouts[homeGoalie.ID] = 0
			}
			if game.AwayScore == 0 {
				goalieShutouts[homeGoalie.ID]++
			}
		}

		// Update shutouts for away goalie
		if awayGoalie != nil {
			if _, ok := goalieShutouts[awayGoalie.ID]; !ok {
				goalieShutouts[awayGoalie.ID] = 0
			}
			if game.HomeScore == 0 {
				goalieShutouts[awayGoalie.ID]++
			}
		}
	}

	// Update goalie wins and shutouts using playerModel methods
	for playerID, wins := range goalieWins {
		player, err := ctrl.playerModel.GetPlayerByID(playerID)
		if err != nil {
			return fmt.Errorf("error getting goalie player with ID %d: %w", playerID, err)
		}
		if _, err := ctrl.playerModel.UpdateGoalieWins(player, wins); err != nil {
			return fmt.Errorf("error updating goalie wins for player with ID %d: %w", playerID, err)
		}
	}

	for playerID, shutouts := range goalieShutouts {
		player, err := ctrl.playerModel.GetPlayerByID(playerID)
		if err != nil {
			return fmt.Errorf("error getting goalie player with ID %d: %w", playerID, err)
		}
		if _, err := ctrl.playerModel.UpdateGoalieShutouts(player, shutouts); err != nil {
			return fmt.Errorf("error updating goalie shutouts for player with ID %d: %w", playerID, err)
		}
	}

	return nil
}
