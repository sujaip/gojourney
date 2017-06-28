package main

import (
	"fmt"
	"os"
)

func main() {
	var numberOfBytes int
	var theError error

	if numberOfBytes, theError = fmt.Printf("Hello, World!\n"); theError != nil {
		os.Exit(1)
	}

	/* This will work as expected as in this case numberChars
	and theError are not scoped to the if block */
	fmt.Printf("Number of bytes printed: %d\n", numberOfBytes)

}
