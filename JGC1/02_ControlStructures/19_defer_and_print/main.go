package main

import "fmt"

func printer(msg string) error {
	/* defer works in first in last out mode
	here the newline is printed last */
	defer fmt.Printf("\n")
	defer fmt.Printf("More\n")
	_, err := fmt.Printf("%s", msg)
	return err
}

func main() {
	printer("Hello, World!\n")
}
