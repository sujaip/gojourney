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
	/* Currently has nothing in it, but can hold up to 4 elements */
	fmt.Printf("%d %d\n", len(words), cap(words)) // 0 4
	words = append(words, "The")
	words = append(words, "Quick")
	words = append(words, "Brown")
	words = append(words, "Fox")
	fmt.Printf("%d %d\n", len(words), cap(words)) // 4 4
	printer(words)
	words = append(words, "Jumps")
	printer(words)
	fmt.Printf("%d %d\n", len(words), cap(words)) // 5 8
	/* when we apend past the capacity, then the capacity is doubled
	   NOTE: We cant assign beyond the capacity as it will cause error */
}
