package main

import (
	"fmt"
	"math/rand"
)

/* any type that has a Len method and Swap method is a shuffler */
type shuffler interface {
	Len() int
	Swap(i, j int)
}

/* shuffle function takes argument of type shuffler interface */
func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type stringSlice []string

func (ss stringSlice) Len() int {
	return len(ss)
}

func (ss stringSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {
	is1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	shuffle(intSlice(is1))
	fmt.Printf("%v\n", is1)

	is2 := intSlice{2, 4, 6, 8, 10, 12}
	shuffle(is2)
	fmt.Printf("%v\n", is2)

	ss := stringSlice{"Sujai", "Sai", "Sanvi"}
	shuffle(ss)
	fmt.Printf("%v\n", ss)

}
