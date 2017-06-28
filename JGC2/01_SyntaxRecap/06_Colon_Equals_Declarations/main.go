package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3, 4, 5, 6}
	stringSlice := []string{"The", "Blue", "Sky"}

	if n := len(intSlice); n > 0 {
		fmt.Printf("intSlice has %d elements\n", n)
	} else {
		fmt.Printf("intSlice has no elements\n")
	}

	if n := len(stringSlice); n > 0 {
		fmt.Printf("stringSlice has %d elements\n", n)
	} else {
		fmt.Printf("stringSlice has no elements\n")
	}
}
