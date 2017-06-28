/*Empty Interface [ interface{} ]
  Interface that does not have any method mentioned in it.
  It is satisfied by every type because there is nothing to satisfy
  It gives us a way to have a unknown type that is determined at runtime
*/
package main

import "fmt"

/* everything satisfies an empty interface */
func whatIsThis(i interface{}) {
	fmt.Printf("%T\n", i)
}

func main() {
	whatIsThis(uint32(42))
	whatIsThis(int(42))
	whatIsThis("Sujai")
}
