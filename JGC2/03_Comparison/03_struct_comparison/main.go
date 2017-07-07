package main

import "fmt"

type MyString struct {
	Str string
}

func main() {
	s := MyString{Str: "Hello"}
	t := MyString{Str: "World!"}

	fmt.Printf("%t\n", s == t)
}

/* Structs can be compared if all the types within them can be compared
including different structs that have the same element types in the same order */
