// Package reksai is a Go wrapper for the League of Legends API endpoints whose full documentation can be found at https://developer.riotgames.com/api-methods/.
//
// Reksai requires an api key to be set as an environment variable API_KEY on your server as it uses os.Getenv() to register your api key as a constant
// for all requests. If you are reciving 403 Forbidden errors it is most likely due to an incorrect or expired key, where as a 401 is most likely
// due to a missing api key.
package reksai

// Version is the currently supported Riot Games API Version for package reksai.
const Version string = "v3"
