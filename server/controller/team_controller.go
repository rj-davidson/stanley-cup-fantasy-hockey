package controller

import (
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

func fetchNHLTeams() ([]NHLTeam, error) {
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

type NHLController struct {
	teamModel *model.TeamModel
}

func NewNHLController(client *ent.Client) *NHLController {
	return &NHLController{
		teamModel: model.NewTeamModel(client),
	}
}

func (ctrl *NHLController) AddNHLTeams() error {
	nhlTeams, err := fetchNHLTeams()
	if err != nil {
		return err
	}

	loadedTeamIDs, err := ctrl.teamModel.ListTeamIDs()

	for _, nhlTeam := range nhlTeams {
		if utils.Contains(loadedTeamIDs, nhlTeam.ID) {
			fmt.Println("Team already added: %s\n", nhlTeam.Name)
			continue
		} else {
			_, err := ctrl.teamModel.CreateTeam(nhlTeam.Name, fmt.Sprintf("nhl_%d_logo.png", nhlTeam.ID), nhlTeam.ID, false)
			if err != nil {
				return err
			} else {
				fmt.Printf("Added NHL team: %s\n", nhlTeam.Name)
			}
		}
	}

	return nil
}
