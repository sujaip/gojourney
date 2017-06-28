package main

import "fmt"

/* Type assertion provides access to an interface value's underlying concrete value

t := i.(T)
This statement asserts that the interface value i holds the concrete type T and
assigns the underlying T value to the variable t

t, ok := i.(T)
If i holds a T, then t will be the underlying value and ok will be true.
If not, ok will be false and t will be the zero value of type T, and no panic occurs.

https://golang.org/doc/effective_go.html#interface_conversions

*/

func whatIsThis(i interface{}) {
	switch val := i.(type) {
	case string:
		fmt.Printf("It's a string %s\n", val)
	case uint32:
		fmt.Printf("Its an unsigned 32-bit integer %d\n", val)
	default:
		fmt.Printf("Don't Know! %v\n", val)
	}
}

func main() {
	whatIsThis(uint32(42))
	whatIsThis(int(-99))
	whatIsThis("Sujai")
	whatIsThis([...]string{"A", "B", "C"})
}
