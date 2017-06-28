package main

import "fmt"

/* slice of int */
type summable []int

/* slices and maps are passed by reference by default */
func (s summable) sum() int {
	sum := 0
	for i, val := range s {
		sum += val
		/* increment each element by one */
		s[i] = val + 1
	}
	return sum
}

func main() {
	s := summable{1, 2, 3, 4, 5}
	fmt.Printf("Sum: %d\n", s.sum())
	fmt.Printf("Sum: %d\n", s.sum())
	fmt.Printf("Sum: %d\n", s.sum())
}
