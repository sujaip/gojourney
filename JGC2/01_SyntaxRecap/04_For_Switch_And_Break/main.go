package main

import (
	"errors"
	"fmt"
)

func main() {
	var even, odd, zeros, total int
	numbers := []int{1, 2, 3, 5, 0, 6, 7, -9}
	var err error

Abort:
	for _, n := range numbers {
		total++
		switch {
		case n == 0:
			zeros++
		case n%2 == 0:
			even++
		case n%2 == 1:
			odd++
		default:
			err = errors.New("Negative Number")
			break Abort
		}
	}

	fmt.Printf("Even %d, Odd %d, Zeros %d, Total %d\n", even, odd, zeros, total)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}
