package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3, 4, 5, 6}
	stringSlice := []string{"The", "Blue", "Sky"}

	n1, err := fmt.Printf("length of intslice %d\n", len(intSlice))
	n2, err := fmt.Printf("length of stringSlice %d\n", len(stringSlice))

	_, err := fmt.Printf("Error: no new variables on left side of :=")

	fmt.Printf("%d, %d, %s", n1, n2, err)

}
