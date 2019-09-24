// Package cheezewizards is an unofficial api client for the Cheeze Wizards' Alchemy API
// https://docs.alchemyapi.io/docs/cheeze-wizards-api
package cheezewizards

import (
	"encoding/json"
	"errors"
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
	fmt.Printf("\nfetching wizard with id=%d", id)

	url := cw.GetBaseURL() + "/wizards/" + strconv.Itoa(id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	cw.setHeaders(req)

	res, err := cw.client.Do(req)
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

	return nil, nil
}

// AttrArgs is the argument struct for GetWizardsByAttributes
// Each attribute is an optional query parameter: affinity, min/max power, and owner
type AttrArgs struct {
	MinPower string // wizards whose current power is greater than or equal to minPower
	MaxPower string // wizards whose current power is less than or equal to maxPower
	Owner    string // wizards owned by this address
	Affinity *int   // wizards with this affinity: 0 = NOTSET, 1 = NEUTRAL, 2 = FIRE, 3 = WIND, 4 = WATER
}

// GetWizardsByAttributes finds wizards by affinity, power, and/or owner
func (cw *CheezeWizards) GetWizardsByAttributes(args AttrArgs) (wizards *[]Wizard, err error) {
	fmt.Printf("\nfetching wizards by args=%+v", args)

	queryParams := "?"
	if args.Owner != "" {
		queryParams = queryParams + "owner=" + args.Owner + "&"
	}
	if args.Affinity != nil {
		queryParams = queryParams + "affinity=" + strconv.Itoa(*args.Affinity) + "&"
	}
	if args.MinPower != "" {
		queryParams = queryParams + "minPower=" + args.MinPower + "&"
	}
	if args.MaxPower != "" {
		queryParams = queryParams + "maxPower=" + args.MaxPower + "&"
	}

	url := cw.GetBaseURL() + "/wizards" + queryParams

	fmt.Println("url=", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	cw.setHeaders(req)

	res, err := cw.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, err
		}

		fmt.Println("body result=", string(body))

		wizards = &[]Wizard{}
		if err := json.Unmarshal(body, wizards); err != nil {
			return nil, err
		}

		fmt.Printf("\nsuccessfully fetched wizards=%+v", wizards)

		return wizards, nil
	}

	return nil, errors.New("unsuccessful response. status=" + res.Status)

}

func (cw *CheezeWizards) setHeaders(req *http.Request) {
	req.Header.Set("x-api-token", cw.key)
	req.Header.Set("x-email", cw.email)
}
