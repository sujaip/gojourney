package main

import "fmt"

/* Variadic function : This function takes one or more string inputs */
func printer(format string, msgs ...string) {
	for _, msg := range msgs {
		fmt.Printf(format, msg)
	}
}

func main() {
	/* Print out the hex (base 16) representation */
	printer("%x\n", "Hello, World!", "From Golang")

	/* Print out the string representation */
	printer("%s\n", "Hello, World!", "From Golang")
}
