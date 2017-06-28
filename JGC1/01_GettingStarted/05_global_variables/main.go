package main

import (
	"fmt"
)

var (
	message = "The answer to life is %d\n"
	answer  = 69
)

/* Here we have used var keyword to declare two global variables
   their values can be changes as they are variables*/

func main() {
	answer++
	fmt.Printf(message, answer)
}
