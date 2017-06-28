package main

import (
	"errors"
	"fmt"
	"os"
)

/* godoc errors New */
var (
	errorEmptyString = errors.New("Can't print an empty string")
)

func printer(msg string) error {
	if msg == "" {
		return errorEmptyString
	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := printer(""); err != nil {
		/* we can compare the error without knowing it's actual string value */
		if err == errorEmptyString {
			fmt.Printf("You tried to print an empty string!\n")
		} else {
			fmt.Printf("Printer Failed: %s\n", err)
		}
		os.Exit(1)
	}
}
