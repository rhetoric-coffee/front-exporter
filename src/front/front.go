package front

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"encoding/json"
	"log"
	"github.com/pkg/errors"
)

const (
	frontApiTokenEnvironmentVariableName = "FRONT_API_TOKEN"
)

type FrontApiMe struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type FrontApiTeam struct {
	_links struct {
		Self string `json:"self"`
	}
	Id   string `json:"id"`
	Name string `json:"name"`
}

type FrontApiTeamsResponse struct {
	_results []FrontApiTeam
}

type FrontApi struct {
	client http.Client
}

func New() (*FrontApi, error) {
	f := FrontApi{}
	f.client.Timeout = time.Second * 10

	// Attempt to connect to the Front API with a token from the environment
	token := os.Getenv(frontApiTokenEnvironmentVariableName)
	meRequest, err := http.NewRequest("GET", "https://api2.frontapp.com/me", nil)
	if err != nil {
		return nil, err
	}
	meRequest.Header.Add("Accept", `application/json`)
	meRequest.Header.Add("Authorization", `Bearer: `+token)
	meReponse, err := f.client.Do(meRequest)
	if err != nil {
		return nil, err
	}
	defer meReponse.Body.Close()
	teamsBody, err := ioutil.ReadAll(meReponse.Body)
	if err != nil {
		return nil, err
	}
	var me FrontApiMe
	if err = json.Unmarshal(teamsBody, &me); err != nil {
		return nil, err
	}
    log.Println("Hit the Front API as "+me.Name)
    return &f, nil
}

func (f FrontApi) ListTeams() (*[]FrontApiTeam, error)  {
    return nil, errors.New("not implemented")
}
