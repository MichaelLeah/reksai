package reksai

import (
	"testing"

	"github.com/michaelleah/reksai"
	"github.com/nbio/st"
	gock "gopkg.in/h2non/gock.v1"
)

func TestByNameWithValidResponse(t *testing.T) {
	defer gock.Off()
	gock.New("https://euw1.api.riotgames.com").
		Get("/lol/summoner/v3/summoners/by-name/pigeon").
		Reply(200).
		JSON(map[string]interface{}{"profileIconId": 1670, "name": "Pigeon", "summonerLevel": 30, "accountId": 28994723, "id": 24537198, "revisionDate": 1500117684000})

	s, err := reksai.Summoner{}.ByName("pigeon", "EUW")

	st.Expect(t, err, nil)
	st.Expect(t, s, &reksai.Summoner{ID: 24537198, AccountID: 28994723, ProfileIconID: 1670, Level: 30, Name: "Pigeon", RevisionDate: 1500117684000})

	st.Expect(t, gock.IsDone(), true)
}
