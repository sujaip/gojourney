package main

import (
	"fmt"
)

func main() {
	atoz := `the "quick" brown fox jumps over the lazy dog\n`
	/* Anything inside a back quoted string is treated litteraly
	here the "" and \n are treated as the strings and printed out
	backquoted strings are particularly useful when encoding json */

	fmt.Printf("%s\n", atoz)
}
