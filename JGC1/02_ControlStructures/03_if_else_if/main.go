package main

import (
	"fmt"
	"os"
)

func main() {
	numberOfBytes, theError := fmt.Printf("Hello, World!\n")

	if theError != nil {
		os.Exit(1)
	} else if numberOfBytes > 10 {
		fmt.Printf("Number of bytes printed: %d\n", numberOfBytes)
	}
}
