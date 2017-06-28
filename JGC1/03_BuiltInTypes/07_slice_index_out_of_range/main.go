package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
}

func main() {
	words := make([]string, 0, 4) //Currently has nothing in it, but can hold up to 4 elements
	fmt.Printf("%d %d\n", len(words), cap(words))
	words = append(words, "The")
	words = append(words, "Quick")
	words = append(words, "Brown")
	words = append(words, "Fox")
	fmt.Printf("%d %d\n", len(words), cap(words))
	printer(words)
	words[5] = "Jumps"
	/* NOTE: Assigning beyond the current capacity will cause Run Time Error: index out of range
	We can still safely append to the slice which will double the capacity */
	printer(words)
}
