package shouter

import (
	"regexp"
	"testing"
)

/* godoc regexp MatchString */
func TestShout(t *testing.T) {
	testString := "The quick brown fox"
	noLower := regexp.MustCompile("[a-z]")

	shoutedString := Shout(testString)
	if noLower.MatchString(shoutedString) {
		t.Fatalf("Found uppercase")
	}

}

/* go test -cover
   go test -coverprofile=coverage.out
   go tool cover -func=coverage.out
   go tool cover -html=coverage.out */
