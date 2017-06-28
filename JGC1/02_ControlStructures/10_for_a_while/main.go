package main

import "fmt"

func main() {
	var counter int

	/* Here the for loop acts like while loop in other languages */
	for counter < 4 {
		fmt.Printf("Hello, World!\n")
		counter++
	}
}
