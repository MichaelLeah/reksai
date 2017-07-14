// Package reksai is a Go wrapper for the League of Legends API endpoints (https://developer.riotgames.com/api-methods/).
//
// reksai requires your api key to be set as an environment variable on your server as it uses os.Getenv() to register your API key as a constant
// for all requests. If you are reciving 403 Forbidden errors it is most likely due to an incorrect, or expired key, where was 401 is most likely
// due to a missing api key.
package reksai
