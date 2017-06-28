package main

import (
	"fmt"
	"os"
)

func main() {
	/* Here we have ignored the number of bytes printed by assigning it to _ */
	_, theError := fmt.Printf("Hello, World!\n")

	if theError != nil {
		os.Exit(1)
	}
}
