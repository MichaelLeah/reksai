package reksai

import (
	"fmt"
	"net/http"
	"time"
)

// RiotAPI handles all requests and interactions with the Riot Games API.
type RiotAPI struct {
	apiKey string
}

// SetAPIKey sets the ApiKey for future requests to use.
func (riot *RiotAPI) SetAPIKey(key string) {
	riot.apiKey = key
}

// Get fires a GET Request to the Riot Games API.
func (riot *RiotAPI) Get(endpoint string) (response *http.Response, err error) {
	if riot.hasAPIKey() == false {
		return nil, fmt.Errorf("api key must be set with SetAPIKey() prior to making a request")
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create new client for request: %v", err)
	}

	req.Header.Set("X-Riot-Token", riot.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to complete request to endpoint: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to api failed with: %v", resp.Status)
	}

	return resp, nil
}

// hasAPIKey checks whether or not the apiKey field is set and returns
// a boolean dependent on the result.
func (riot *RiotAPI) hasAPIKey() bool {
	if riot.apiKey == "" {
		return false
	}

	return true
}
