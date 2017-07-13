package reksai

// Summoner is a representation of the SummonerDTO returned from the API.
type Summoner struct {
	ID            int64  `json:"id"`
	AccountID     int64  `json:"accountId"`
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	Level         string `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
}

// ByName hits /summoners/by-name/{summoner-name} endpoint and returns a
// summoner struct instance.
//  reksai.Summoner{}.ByName("Pigeon", "EUW")
func (s Summoner) ByName(name string, region string) Summoner {
	return s // Todo implement method.
}
