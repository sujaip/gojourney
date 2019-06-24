package main

import "fmt"

func increment(inc int) {
  // Increment the "value of" inc
	inc++
	fmt.Println("inc:\tValue:", inc, "\tAddress:", &inc)
}

func main() {

  count := 10

	fmt.Println("count:\tValue:", count, "\tAddress:", &count)

	// Pass the "value of" count.
	increment(count)

	fmt.Println("count:\tValue:", count, "\tAddress:", &count)
}
