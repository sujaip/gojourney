package main

import (
	"fmt"
)

func main() {
	// var message string
	// message = "Hello, World!\n"

	/* Here := is used for both setting a value and declaring the variable type*/
	/* Go figures out the type by looking at the right hand side*/
	message := "Hello, World!\n"
	fmt.Printf(message)
}

/* We can get help on fmt package from golang.org/pkg/fmt
   or documentation on specific functions by running `go doc fmt Printf` */
