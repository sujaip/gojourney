package main

import "fmt"

/* []string here means slice of strings */
func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
}

func main() {
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	/* Slice of slice */
	printer(words[:])   //Print everything
	printer(words[:2])  //Two elements from the begining
	printer(words[5:])  //from fifth element to last element
	printer(words[2:4]) //Item two and item three
	words[2] = "blue"
	printer(words[2:4]) // blue fox

}
