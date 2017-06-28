/*Empty Interface [ interface{} ]
  Interface that does not have any method mentioned in it.
  It is satisfied by every type because there is nothing to satisfy
  It gives us a way to have a unknown type that is determined at runtime
*/
package main

import "fmt"

func whatIsThis(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("It's a string\n")
	case uint32:
		fmt.Printf("Its an unsigned 32-bit integer\n")
	default:
		fmt.Printf("Don't Know!\n")
	}
}

func main() {
	whatIsThis(uint32(42))
	whatIsThis(int(42))
	whatIsThis("Sujai")
}
