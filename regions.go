package reksai

// RegionCode is an alias of type string, it represents all of the Service Regions used by Riot Games.
type RegionCode string

// Region is a representation of the required region information needed for a request, including a SecondaryPlatform for calls
// to NA which have two regional platforms: NA and NA1.
type Region struct {
	Code              string
	Platform          string
	SecondaryPlatform string
	Host              string
}

// Empty checks whether a Region struct has any content
func (r Region) Empty() bool {
	return r == *new(Region)
}

// Regions maps a given region code to the necessary request information for that service platform.
var Regions = map[RegionCode]Region{
	"BR": Region{
		Code:              "BR",
		Platform:          "BR1",
		SecondaryPlatform: "",
		Host:              "br1.api.riotgames.com",
	},
	"EUNE": Region{
		Code:              "EUNE",
		Platform:          "EUN1",
		SecondaryPlatform: "",
		Host:              "eun1.api.riotgames.com",
	},
	"EUW": Region{
		Code:              "EUW",
		Platform:          "EUW1",
		SecondaryPlatform: "",
		Host:              "euw1.api.riotgames.com",
	},
	"JP": Region{
		Code:              "JP",
		Platform:          "JP1",
		SecondaryPlatform: "",
		Host:              "jp1.api.riotgames.com",
	},
	"KR": Region{
		Code:              "KR",
		Platform:          "KR",
		SecondaryPlatform: "",
		Host:              "kr.api.riotgames.com",
	},
	"LAN": Region{
		Code:              "LAN",
		Platform:          "LA1",
		SecondaryPlatform: "",
		Host:              "la1.api.riotgames.com",
	},
	"LAS": Region{
		Code:              "LAS",
		Platform:          "LA2",
		SecondaryPlatform: "",
		Host:              "la2.api.riotgames.com",
	},
	"NA": Region{
		Code:              "NA",
		Platform:          "NA1",
		SecondaryPlatform: "NA",
		Host:              "na1.api.riotgames.com",
	},
	"OCE": Region{
		Code:              "OCE",
		Platform:          "OC1",
		SecondaryPlatform: "",
		Host:              "oc1.api.riotgames.com",
	},
	"TR": Region{
		Code:              "TR",
		Platform:          "TR1",
		SecondaryPlatform: "",
		Host:              "tr1.api.riotgames.com",
	},
	"RU": Region{
		Code:              "RU",
		Platform:          "RU1",
		SecondaryPlatform: "",
		Host:              "ru.api.riotgames.com",
	},
	"PBE": Region{
		Code:              "PBE",
		Platform:          "PBE1",
		SecondaryPlatform: "",
		Host:              "PBE1",
	},
}
