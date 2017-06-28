package main

import (
	"fmt"
	"os"
)

func printer(msg string) error {

	/* godoc os Create */
	f, err := os.Create("HelloWorld.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(msg)
	return err
}

func main() {
	err := printer("Hello, World!")
	if err == nil {
		fmt.Printf("Success")
	}
}
