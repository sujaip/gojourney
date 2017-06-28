package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
}

func main() {
	words := make([]string, 0, 4)
	fmt.Printf("%d %d\n", len(words), cap(words)) // 0 4
	words = append(words, "The")
	words = append(words, "Quick")
	words = append(words, "Brown")
	words = append(words, "Fox")
	fmt.Printf("%d %d\n", len(words), cap(words)) // 4 4
	printer(words)

	newWords := make([]string, 4)
	copy(newWords, words)
	fmt.Printf("%d %d\n", len(newWords), cap(newWords)) // 4 4
	newWords[2] = "Blue"                                // Just the copy changes, original still remains same
	printer(newWords)                                   // The Quick Blue Fox
	printer(words)                                      // The Quick Brown Fox
}
