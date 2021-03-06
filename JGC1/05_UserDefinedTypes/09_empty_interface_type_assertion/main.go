package main

import "fmt"

func whatIsThis(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("It's a string %s\n", i.(string))
	case uint32:
		fmt.Printf("Its an unsigned 32-bit integer %d\n", i.(uint32))
	default:
		fmt.Printf("Don't Know!\n")
	}
}

func main() {
	whatIsThis(uint32(42))
	whatIsThis(int(-99))
	whatIsThis("Sujai")
	whatIsThis([...]string{"A", "B", "C"})
}
