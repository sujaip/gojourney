package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	people := make(map[Person]struct{})
	/* in the above statement struct{} denotes the type empty anonymous struct */

	people[Person{"Sujai"}] = struct{}{}
	/* in the above statement struct{}{} denotes the value empty anonymous struct */

	people[Person{"Sanvi"}] = struct{}{}
	fmt.Printf("%#v\n", people)

	_, ok := people[Person{"Sujai"}]
	fmt.Printf("%t\n", ok)
}
