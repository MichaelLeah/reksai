package reksai

// Summoner is a representation of the SummonerDTO returned from the api.
type Summoner struct {
	ID            int64  `json:"id"`
	AccountID     int64  `json:"accountId"`
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	Level         int    `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
}

// ByName ...
// func (s Summoner) ByName(name string, region RegionCode) (summoner *Summoner, err error) {
// 	endpoint := fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/%s/summoners/by-name/%s", regions[region], Version, name)

// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", endpoint, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to create new client for request: %v", err)
// 	}

// 	req.Header.Set("X-Riot-Token", os.Getenv("API_KEY"))

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to complete request to endpoint: %v", err)
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		return nil, fmt.Errorf("request to api failed with: %v", resp.Status)
// 	}

// 	respBody, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to read response body: %v", err)
// 	}

// 	if err := json.Unmarshal([]byte(respBody), &summoner); err != nil {
// 		return nil, fmt.Errorf("unable to unmarshal response body to summoner struct: %v", err)
// 	}

// 	return summoner, nil
// }
