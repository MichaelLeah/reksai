package reksai

import (
	"os"
)

// VERSION is the current version of the API that this wrapper supports.
const VERSION = "v3"

// API_KEY is a required environment variable required for all calls to the Riot Games API.
var API_KEY = os.Getenv("API_KEY")

// REGIONS acts to map a given region to their respective domain for use in building up request endpoints.
var REGIONS = map[string]string{
	"BR":   "br1",
	"EUNE": "eun1",
	"EUW":  "euw1",
	"JP":   "jp1",
	"KR":   "kr",
	"LAN":  "la1",
	"LAS":  "la2",
	"NA":   "na1",
	"OCE":  "oc1",
	"RU":   "ru",
	"PBE":  "pbe1",
}
