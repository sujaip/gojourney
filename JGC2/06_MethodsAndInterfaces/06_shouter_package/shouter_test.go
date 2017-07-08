package shouter

import (
	"regexp"
	"testing"
)

func TestShout(t *testing.T) {
	testString := "The quick brown fox"
	noLower := regexp.MustCompile("[a-z]")

	shoutedstring := Shout(testString)
	if noLower.MatchString(shoutedstring) {
		t.Fatalf("Found lowercase")
	}
}

/*
Now that we have a new method "ReadAndShout" in shouter package that reads from a file and shouts
We should create a temporary file and pass it in case we want to test the "ReadAndShout" method

*********
If we change ReadAndShout method to accept an interface instead of a file,
then we can avoid creating a temporary file for testing purpose
*********

*/
