package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World!\n")

	/* There is no break statement in the below switch case
	   because switch case does not fallthrough in Go */
	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("No bytes printed\n")
	case n != 14:
		fmt.Printf("Wrong number of characters: %d\n", n)
	default:
		fmt.Printf("OK!")
	}

	fmt.Printf("\n")

}
