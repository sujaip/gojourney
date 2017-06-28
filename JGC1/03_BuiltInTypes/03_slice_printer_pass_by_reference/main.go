package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")

	w[2] = "blue"
}

/* https://blog.golang.org/go-slices-usage-and-internals */

func main() {
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	printer(words)
	/* Slices are passed by reference, and hence the value is changed */
	printer(words)

	/* Type specification for a slice is []T, where T is the type of the elements of the slice.
		   Unlike an array type, a slice type has no specified length
			 A slice literal is declared just like an array literal, except you leave out the element count:

	    letters := []string{"a", "b", "c", "d"} */

	/* A slice can be created with the built-in function called make, which has the signature,
	  	func make([]T, len, cap) []T
		  where T stands for the element type of the slice to be created.
			The make function takes a type, a length, and an optional capacity.
			When called, make allocates an array and returns a slice that refers to that array. */

}
