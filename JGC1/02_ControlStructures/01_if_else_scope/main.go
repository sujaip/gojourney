package main

import (
	"fmt"
	"os"
)

func main() {
	/* Here numberChars and theError are scoped to the if else block */
	if numberOfBytes, theError := fmt.Printf("Hello, World!\n"); theError != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Number of bytes printed: %d\n", numberOfBytes)
	}

	/* godoc fmt Printf
	   godoc os Exit */

}
