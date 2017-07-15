package reksai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Summoner is a representation of the SummonerDTO returned from the api.
type Summoner struct {
	ID            int64  `json:"id"`
	AccountID     int64  `json:"accountId"`
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	Level         int    `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
}

// ByName queries the /summoners/by-name/{name} endpoint and attempts to create a Summoner struct instance
// from the result. Failure at any point will return an error to be handle.
func (s Summoner) ByName(name string, region RegionCode) (summoner *Summoner, err error) {
	r := Regions[region]
	if r.Empty() {
		return nil, fmt.Errorf("unable to get region host information for: %v", region)
	}

	endpoint := fmt.Sprintf("https://%s/lol/summoner/%s/summoners/by-name/%s", r.Host, Version, name)

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create new client for request: %v", err)
	}

	req.Header.Set("X-Riot-Token", os.Getenv("API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to complete request to endpoint: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request to api failed with: %v", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %v", err)
	}

	if err := json.Unmarshal([]byte(respBody), &summoner); err != nil {
		return nil, fmt.Errorf("unable to unmarshal response body to summoner struct: %v", err)
	}

	return summoner, nil
}
