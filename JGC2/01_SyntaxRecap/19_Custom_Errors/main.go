package main

/* godoc fmt Errorf */

import (
	"errors"
)

var (
	ErrorBadStartup = errors.New("Failed to start properly")
)

func main() {
	err := ErrorBadStartup
	if err == ErrorBadStartup {
		// Ignore and return
		return
	}
}

/* Comparison of errors with == or != is a very good stuff */
