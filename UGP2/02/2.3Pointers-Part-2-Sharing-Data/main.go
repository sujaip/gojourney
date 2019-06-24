package main

import "fmt"

func increment(inc *int) {
  // Increment the "value of" count that the "pointer points to"
	*inc++
	fmt.Println("inc:\tValue:", inc, "\tAddress:", &inc, "\tValue Points To:", *inc)
}

func main() {

  count := 10

	fmt.Println("count:\tValue:", count, "\tAddress:", &count)

	// Pass the "address of" count.
	increment(&count)

	fmt.Println("count:\tValue:", count, "\tAddress:", &count)
}
