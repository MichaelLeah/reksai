package reksai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// ByName hits /summoners/by-name/{summoner-name} endpoint and maps the response to a summoner struct instance on success.
// If there is any issue with the request the method will return an error.
//
// You can call this method like so:
//  reksai.Summoner{}.ByName("Pigeon", "EUW")
func (s Summoner) ByName(name string, region string) (summoner *Summoner, err error) {
	endpoint := fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/%s/summoners/by-name/%s", REGIONS[region], VERSION, name)

	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create new client for request: %v", err)
	}

	req.Header.Set("X-Riot-Token", API_KEY)

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
