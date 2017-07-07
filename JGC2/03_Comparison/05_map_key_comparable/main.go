package main

import "fmt"

type Person struct {
	First, Last string
}

func main() {
	people := make(map[Person]bool)
	people[Person{"Sujai", "Prakasam"}] = true
	people[Person{"Sanvi", "Sujai"}] = true
	fmt.Printf("%v\n", people)
	fmt.Printf("%#v\n", people)
}

/*
Anything that is comparable can be used as key in a map.
Slice is not comparable in Go,
   i.e. we can't compare a slice with another slice, even if it points to the same underlying array
*/
