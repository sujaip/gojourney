package main

import (
	"errors"
	"fmt"
	"os"
)

/* https://blog.golang.org/defer-panic-and-recover */

/* godoc errors New */
var (
	errorEmptyString = errors.New("Can't print an empty string")
)

func printer(msg string) error {
	if msg == "" {
		/* Panic instead of returning the error */
		panic(errorEmptyString)
	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := printer(""); err != nil {
		//here we are able to compare the error without knowing it's actual string value
		if err == errorEmptyString {
			fmt.Printf("You tried to print an empty string!\n")
		} else {
			fmt.Printf("Printer Failed: %s\n", err)
		}
		os.Exit(1)
	}
}
