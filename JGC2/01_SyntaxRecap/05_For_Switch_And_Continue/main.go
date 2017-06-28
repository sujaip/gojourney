package main

import (
	"fmt"
)

func main() {
	var even, odd, zeros, total int
	numbers := []int{1, 2, 3, 5, 0, 6, 7, -9}

IgnoreNegative:
	for _, n := range numbers {
		switch {
		case n == 0:
			zeros++
		case n%2 == 0:
			even++
		case n%2 == 1:
			odd++
		default:
			continue IgnoreNegative
		}
		total++
	}

	fmt.Printf("Even %d, Odd %d, Zeros %d, Total %d\n", even, odd, zeros, total)

}
