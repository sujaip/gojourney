package main

import (
	"fmt"
)

func main() {
	isTrue := !true
	/* default value of a declared but uninitialized variable in go is zero value of that type
	example:
	var isTrue bool
	fmt.Printf("Value: %t\n", isTrue)
	will print "Value: false", because false is the zero value of bool type
	*/
	fmt.Printf("Value: %t\n", isTrue)
}
