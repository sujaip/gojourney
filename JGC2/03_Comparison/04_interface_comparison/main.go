package main

import "fmt"

type MyFirstInterface interface {
	String() string
}

type MySecondInterface interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	var i MyFirstInterface
	var j MySecondInterface
	p := Person{"Sujai", 37}
	fmt.Printf("%t\n", i == j)
	fmt.Printf("%t\n", p == j)

}

/* We can compare different interfaces that have the same methods
   we can also compare a concrete type with an interface as in this example

   Key type used in map must be comparable */
