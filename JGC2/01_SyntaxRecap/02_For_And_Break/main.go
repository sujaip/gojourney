package main

import "fmt"

func main() {
	var even, odd int
	numbers := []int{1, 2, 3, 5, 0, 6, 7, 9}

Abort:
	for i := 0; i < 10; i++ {
		for _, n := range numbers {
			if n == 0 {
				break Abort
			}
			if n%2 == 0 {
				even++
			} else {
				odd++
			}
		}
	}
	fmt.Printf("Even %d, Odd %d\n", even, odd)

}
