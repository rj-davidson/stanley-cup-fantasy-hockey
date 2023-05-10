package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rj-davidson/stanley-cup-fantasy-hockey/db/ent"
	"github.com/rj-davidson/stanley-cup-fantasy-hockey/model"
)

type NHLBoxScore struct {
	GameID int `json:"gamePk"`
	Status struct {
		AbstractGameState string `json:"abstractGameState"`
	} `json:"status"`
	Teams struct {
		Away struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
			Score int `json:"score"`
		} `json:"away"`
		Home struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
			Score int `json:"score"`
		} `json:"home"`
	} `json:"teams"`
}

const nhlApiGameUrlFormat = "https://statsapi.web.nhl.com/api/v1/game/%d/boxscore"

func fetchNHLBoxScore(gameID int) (*NHLBoxScore, error) {
	httpClient := &http.Client{Timeout: 10 * time.Second}

	gameUrl := fmt.Sprintf(nhlApiGameUrlFormat, gameID)
	resp, err := httpClient.Get(gameUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse NHLBoxScore
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

type GameController struct {
	gameModel *model.GameModel
}

func NewGameController(client *ent.Client) *GameController {
	return &GameController{
		gameModel: model.NewGameModel(client),
	}
}

func (ctrl *GameController) FetchPostSeasonGames(gameIDs []int) error {
	for _, gameID := range gameIDs {
		nhlBoxScore, err := fetchNHLBoxScore(gameID)
		if err != nil {
			return err
		}

		fmt.Printf(
			"GameID: %d | Status: %s | Away Team: %s (%d) | Home Team: %s (%d)\n",
			nhlBoxScore.GameID,
			nhlBoxScore.Status.AbstractGameState,
			nhlBoxScore.Teams.Away.Team.Name,
			nhlBoxScore.Teams.Away.Score,
			nhlBoxScore.Teams.Home.Team.Name,
			nhlBoxScore.Teams.Home.Score,
		)

		// Add game creation or update logic her
	}

	return nil
}
