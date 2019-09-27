// Package cheezewizards is an unofficial api client for the Cheeze Wizards' Alchemy API
// https://docs.alchemyapi.io/docs/cheeze-wizards-api
package cheezewizards

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// CheezeWizards is the api client
type CheezeWizards struct {
	key    string
	email  string
	client *http.Client
	env    string
}

// Wizard is a single cheeze wizard
type Wizard struct {
	ID                    string `json:"id"`
	Owner                 string `json:"owner"`
	Affinity              int    `json:"affinity"`
	InitialPower          string `json:"initialPower"`
	Power                 string `json:"power"`
	EliminatedBlockNumber *int   `json:"eliminatedBlockNumber"`
	CreatedBlockNumber    int    `json:"createdBlockNumber"`
}

// Duel is a single duel result between to cheeze wizards
type Duel struct {
	ID                string `json:"id"`
	Wizard1ID         string `json:"wizard1Id"`
	Wizard2ID         string `json:"wizard2Id"`
	Affinity1         int    `json:"affinity1"`
	Affinity2         int    `json:"affinity2"`
	StartPower1       string `json:"startPower1"`
	StartPower2       string `json:"startPower2"`
	EndPower1         string `json:"endPower1"`
	EndPower2         string `json:"endPower2"`
	MoveSet1          string `json:"moveSet1"`
	MoveSet2          string `json:"moveSet2"`
	StartBlock        int    `json:"startBlock"`
	EndBlock          int    `json:"endBlock"`
	TimeOutBlock      int    `json:"timeoutBlock"`
	TimedOut          bool   `json:"timedOut"`
	IsAscensionBattle bool   `json:"isAscensionBattle"`
}

// GetBaseURL is the base cheeze wizard url
func (cw *CheezeWizards) GetBaseURL() string {
	if cw.env == "mainnet" {
		return "https://cheezewizards-mainnet.alchemyapi.io"
	}
	if cw.env == "rinkeby" {
		return "https://cheezewizards-rinkeby.alchemyapi.io"
	}
	return "https://cheezewizards.alchemyapi.io"
}

// NewCheezeWizards inits a cheeze wizards API client
func NewCheezeWizards(apiKey string, email string) *CheezeWizards {
	client := http.Client{}

	return &CheezeWizards{key: apiKey, email: email, client: &client}
}

// GetWizardByID finds a wizard by id
func (cw *CheezeWizards) GetWizardByID(id int) (wizard *Wizard, err error) {
	fmt.Printf("\nfetching wizard by id=%d.", id)

	url := cw.GetBaseURL() + "/wizards/" + strconv.Itoa(id)

	res, err := cw.performRequest(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, err
		}

		wizard := &Wizard{}
		if err := json.Unmarshal(body, wizard); err != nil {
			return nil, err
		}

		fmt.Printf("\nsuccessfully fetched wizard=%+v", wizard)

		return wizard, nil

	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return nil, fmt.Errorf("unsuccessful response. status=%s. msg=%s", res.Status, string(body))
}

// GetWizardsByAttributes finds wizards by affinity, power, and/or owner
// owner - wizards owned by this address
// affinity - wizards with this affinity: 0 = NOTSET, 1 = NEUTRAL, 2 = FIRE, 3 = WIND, 4 = WATER
// minPower - wizards whose current power is greater than or equal to minPower
// maxPower - wizards whose current power is less than or equal to maxPower
func (cw *CheezeWizards) GetWizardsByAttributes(owner string, affinity, minPower, maxPower string) (wizards *[]Wizard, err error) {
	fmt.Printf("\nfetching wizards by: owner=%s. affinity=%s. minPower=%s. maxPower=%s.", owner, affinity, minPower, maxPower)

	queryParams := "?"
	if owner != "" {
		queryParams = queryParams + "owner=" + owner + "&"
	}
	if affinity != "" {
		queryParams = queryParams + "affinity=" + affinity + "&"
	}
	if minPower != "" {
		queryParams = queryParams + "minPower=" + minPower + "&"
	}
	if maxPower != "" {
		queryParams = queryParams + "maxPower=" + maxPower
	}

	url := cw.GetBaseURL() + "/wizards" + queryParams

	res, err := cw.performRequest(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, err
		}

		wizards = &[]Wizard{}
		if err := json.Unmarshal(body, wizards); err != nil {
			return nil, err
		}

		fmt.Printf("\nsuccessfully fetched wizards=%+v", wizards)

		return wizards, nil
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return nil, fmt.Errorf("unsuccessful response. status=%s. msg=%s", res.Status, string(body))
}

// GetDuelByID finds a duel by id
func (cw *CheezeWizards) GetDuelByID(id int) (duel *Duel, err error) {
	fmt.Printf("\nfetching duel with id=%d", id)

	url := cw.GetBaseURL() + "/duels/" + strconv.Itoa(id)

	res, err := cw.performRequest(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, err
		}

		duel := &Duel{}
		if err := json.Unmarshal(body, duel); err != nil {
			return nil, err
		}

		fmt.Printf("\nsuccessfully fetched duel=%+v", duel)

		return duel, nil

	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return nil, fmt.Errorf("unsuccessful response. status=%s. msg=%s", res.Status, string(body))
}

// GetDuelsByAttributes fetches a list of duels
// all query params are optional
// wizardIds - duels involving these wizards (coma separated list of wizardIds)
// excludeInProgress - true for completed duels, false for all duels (default)
// excludeFinished - true for duels in progress, false for all duels (default)
// startBlockFrom - duels whose start block is greater than or equal to this block number
// startBlockTo - duels whose start block is less than or equal to this block number
// endBlockFrom - duels whose end block is greater than or equal to this block number
// endBlockTo - duels whose end block is less than or equal to this block number
func (cw *CheezeWizards) GetDuelsByAttributes(wizardIds, excludeInProgress, excludeFinished, startBlockFrom, startBlockTo, endBlockFrom, endBlockTo string) (duels *[]Duel, err error) {
	fmt.Printf(`\nfetching wizards by: wizardIds=%s. excludeInProgress=%s. excludeFinished=%s. startBlockFrom=%s. startBlockTo=%s. endBlockFrom=%s. endBlockTo=%s.`, wizardIds, excludeInProgress, excludeFinished, startBlockFrom, startBlockTo, endBlockFrom, endBlockTo)

	queryParams := "?"
	if wizardIds != "" {
		queryParams = queryParams + "wizardIds=" + wizardIds + "&"
	}
	if excludeInProgress != "" {
		queryParams = queryParams + "excludeInProgress=" + excludeInProgress + "&"
	}
	if excludeFinished != "" {
		queryParams = queryParams + "excludeFinished=" + excludeFinished + "&"
	}
	if startBlockFrom != "" {
		queryParams = queryParams + "startBlockFrom=" + startBlockFrom + "&"
	}
	if startBlockTo != "" {
		queryParams = queryParams + "startBlockTo=" + startBlockTo + "&"
	}
	if endBlockFrom != "" {
		queryParams = queryParams + "endBlockFrom=" + endBlockFrom + "&"
	}
	if endBlockTo != "" {
		queryParams = queryParams + "endBlockTo=" + endBlockTo
	}

	url := cw.GetBaseURL() + "/duels" + queryParams

	res, err := cw.performRequest(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, err
		}

		type apiResponse struct {
			Data []Duel `json:"duels"`
		}

		res := apiResponse{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}

		return &res.Data, nil
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return nil, fmt.Errorf("unsuccessful response. status=%s. msg=%s", res.Status, string(body))
}

// performRequest sets the auth headers and performs the http request
func (cw *CheezeWizards) performRequest(url string) (res *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	cw.setHeaders(req)

	return cw.client.Do(req)
}

func (cw *CheezeWizards) setHeaders(req *http.Request) {
	req.Header.Set("x-api-token", cw.key)
	req.Header.Set("x-email", cw.email)
}
