package main

import (
	"fmt"
	"os"
)

func main() {
	/* Notice the ; after the end of Printf statement
	   This is one of the few places in Go where ; is needed */
	if _, theError := fmt.Printf("Hello, World!\n"); theError != nil {
		os.Exit(1)
	}
}
