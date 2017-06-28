package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("")

	/* Complex conditionls in switch cases
	   We have also explicitly specified fallthrough */
	switch {
	case err != nil:
		os.Exit(1)
	case n == 0 && err == nil:
		fmt.Printf("Empty String\n")
		fallthrough
	case n != 15:
		fmt.Printf("Wrong number of characters: %d\n", n)
		fallthrough
	default:
		fmt.Printf("OK!")
	}

	fmt.Printf("\n")

}
