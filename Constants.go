package reksai

import (
	"os"
)

// API_KEY is required for all calls to the Riot Games API
var API_KEY = os.Getenv("API_KEY")

// VERSION is the current version of the API that this wrapper supports
const VERSION = "v3"

const (
	// BR is the regional end point for Brazil
	BR = "br1"

	// EUNE is the regional end point for Europe North East
	EUNE = "eun1"

	// EUW is the regional end point for Europe West
	EUW = "euw1"

	// JP is the regional end point for Japan
	JP = "jp1"

	// KR is the regional end point for Korea
	KR = "kr"

	// LAN is the regional end point for Latin America North
	LAN = "la1"

	// LAS is the regional end point for Latin America South
	LAS = "la2"

	// NA is the regional end point for North America
	NA = "na1"

	// OCE is the regional end point for Oceania
	OCE = "oc1"

	// TR is the end point for Turkey
	TR = "tr1"

	// RU is the regional end point for Russia
	RU = "ru"

	// PBE is the regional end point for the Public Test Realms
	PBE = "pbe1"
)
