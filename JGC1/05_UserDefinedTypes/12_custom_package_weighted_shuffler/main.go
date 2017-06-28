package main

import (
	"fmt"
	"shuffler"
)

type weightedString struct {
	weight int
	s      string
}

type sliceOfStructs []weightedString

func (is sliceOfStructs) Len() int {
	return len(is)
}

func (is sliceOfStructs) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is sliceOfStructs) Weight(i int) int {
	return is[i].weight
}

func main() {
	/* Slice of structs */
	i := sliceOfStructs{weightedString{100, "Hello"}, weightedString{200, "World!"}, weightedString{10, "Good Bye"}}
	fmt.Printf("%v\n", i)
	shuffler.WeightedShuffle(i)
	fmt.Printf("%v\n", i)
	fmt.Printf("Counter: %d\n", shuffler.GetCount())
}
