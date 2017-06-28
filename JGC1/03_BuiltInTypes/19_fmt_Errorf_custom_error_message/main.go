package main

import (
	"fmt"
	"os"
)

/* godoc fmt Errorf */
func printer(msg string) error {
	if msg == "" {
		/* Errorf formats a string according to Printf's rules and
		   returns it as an error created by errors.New */
		return fmt.Errorf("Can't print an empty string")
	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := printer(""); err != nil {
		fmt.Printf("Printer Failed: %s\n", err)
		os.Exit(1)
	}
}
