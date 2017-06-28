package main

import "fmt"

func main() {
	var even, odd, zeros, total int
	numbers := []int{1, 2, 3, 5, 0, 6, 7, 9}

	for _, n := range numbers {
		total++
		switch {
		case n == 0:
			zeros++
		case n%2 == 0:
			even++
		default:
			odd++
		}
	}
	fmt.Printf("Even %d, Odd %d, Zeros %d, Total %d\n", even, odd, zeros, total)

}
