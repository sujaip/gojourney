package main

import "fmt"

/* In Go the size of an arry is part of the type,
like [16]int, and pointer to that would be *[16]int */
func realPrinter(w *[4]string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")

	w[2] = "blue"
}

func printer(w [4]string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")

	w[2] = "black"
}

func main() {
	words := [4]string{"the", "quick", "brown", "fox"}
	realPrinter(&words)

	printer(words)
	printer(words)
	/* Arrays in go are passed-by-value by default (Copy)
	unless explicitly specified to use pointers */
}
