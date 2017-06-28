package main

import "fmt"

func main() {
	/* [...]string means an array of unknown number of strings
	This can also be mentioned as words := [4]string{"the", "quick", "brown", "fox"} */
	words := [...]string{"the", "quick", "brown", "fox"}
	fmt.Printf("%s %d\n", words[2], len(words))

	/* In Go the size of an arry is part of the type,
	like [16]int, and pointer to that would be *[16]int */

	/* An array literal can be specified like so:
	 b := [2]string{"Penn", "Teller"}

	Or, you can have the compiler count the array elements for you:
	b := [...]string{"Penn", "Teller"} */

}
