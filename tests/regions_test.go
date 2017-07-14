package reksai

import (
	"testing"

	"github.com/michaelleah/reksai"
)

type testSuite struct {
	Input          reksai.RegionCode
	ExpectedResult bool
}

var testCases = []testSuite{
	{"EUW", false},
	{"EU", true},
	{"", true},
	{"世界", true},
}

func TestEmpty(t *testing.T) {
	for _, test := range testCases {
		region := reksai.Regions[test.Input]
		if result := region.Empty(); result != test.ExpectedResult {
			t.Error("for case:", test.Input, "expected:", test.ExpectedResult, "got:", result)
		}
	}
}
