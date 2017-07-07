package main

import "fmt"

func main() {
	var i int = 42
	var j int8 = 42
	fmt.Printf("%t\n", i == j)
	/* invalid operation: i == j (mismatched types int and int8) */
}
