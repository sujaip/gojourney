package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
}

func main() {
	/* make function allocates a zeroed array and returns a slice that refers to that array */
	words := make([]string, 4) /* len(words) = 4*/
	words[0] = "The"
	words[1] = "Quick"
	words[2] = "Brown"
	words[3] = "Fox"
	printer(words)

	/* we can specify the capacity of the slice by passing a third argument to make
	   b := make([]int, 0, 5)
	   Here len(b) = 0, and cap(b) = 5 */

}
