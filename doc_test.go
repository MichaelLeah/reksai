package reksai

import (
	"fmt"
)

func ExampleSummoner_ByName() {
	summoner, err := Summoner{}.ByName("pigeon", "EUW")
	if err != nil {
		// Handle error
	}
	fmt.Printf("%T loaded with values: %v", summoner, summoner)
	// Output: *reksai.Summoner loaded with values: &{24537198 28994723 1670 Pigeon 30 1500117684000}
}
