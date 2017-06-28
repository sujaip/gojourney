package main

import (
	"fmt"
	"shuffler"
)

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	i := intSlice{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("%v\n", i)
	shuffler.Shuffle(i)
	fmt.Printf("%v\n", i)
}
