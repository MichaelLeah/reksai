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
