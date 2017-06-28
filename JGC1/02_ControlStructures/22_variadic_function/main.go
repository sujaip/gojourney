package main

import "fmt"

/* Variadic function : This function takes one or more string inputs */
func printer(msgs ...string) {
	/* range over the messages */
	for _, msg := range msgs {
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	printer("Hello, World!", "From Golang", "by Sujai")
}
