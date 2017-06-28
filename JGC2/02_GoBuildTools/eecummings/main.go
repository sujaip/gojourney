package eecummings

import "strings"
import "regexp"

var noPuncs = regexp.MustCompile("[^a-z ]")

func Simplify(s string) string {
	return noPuncs.ReplaceAllString(strings.ToLower(s), "")
}
