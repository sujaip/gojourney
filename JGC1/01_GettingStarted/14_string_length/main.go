package main

import (
	"fmt"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	// Here the _ is optional, second parameter can be ommited
	for i, _ := range atoz {
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("Length of the string: %d\n", len(atoz))
	/* here the length of the string len(atoz) will be one more
	than the last index as the index begins from 0*/
}
