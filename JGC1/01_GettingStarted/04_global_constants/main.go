package main

import (
	"fmt"
)

const (
	message = "The answer to life is %d\n"
	answer  = 69
)

/* Here we don't need to use the short assignment operator :=
because we are declaring the constant using const keyword */

func main() {
	fmt.Printf(message, answer)
}
