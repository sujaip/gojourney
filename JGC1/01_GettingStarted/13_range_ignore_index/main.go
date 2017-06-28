package main

import (
	"fmt"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	/* we can omit the index by assigning it to _ as below*/
	for _, char := range atoz {
		fmt.Printf("%c\n", char)
	}

	/* we could also use only the index like below
	for i := range atoz {
		fmt.Printf("%d\n", i)
	}
	here the second value (char) is optional and need not be assigned */
}
