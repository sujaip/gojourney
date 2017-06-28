package main

import (
	"fmt"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	/* here i is index and char is the character at that index*/
	for i, char := range atoz {
		fmt.Printf("%d %c\n", i, char)
	}
}
