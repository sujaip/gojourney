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

	newWords := words
	fmt.Printf("%d %d\n", len(newWords), cap(newWords))
	newWords[2] = "Blue"
	printer(newWords) // The Quick Blue Fox
	printer(words)    // The Quick Blue Fox
	/* Here both newWords and words refer to the same underlying array
	i.e. newWords is deep copy of words */
}
