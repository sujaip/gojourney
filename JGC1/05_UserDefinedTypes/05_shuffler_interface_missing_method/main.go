package main

import (
	"fmt"
	"math/rand"
)

/* Interfaces */
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

/* In Go the interface compliance checks are made at compile time
   Here we have renamed the Len function to Lenghth, and don't satisfy the shuffler interface
   We should get a compile time error about the same like below
   "cannot use ss (type stringSlice) as type shuffler in argument to shuffle:
	 stringSlice does not implement shuffler (missing Len method)"
*/
func (ss stringSlice) Length() int {
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
