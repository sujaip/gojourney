package main

import "fmt"

func printer(msg string) error {
	/* godoc fmt Printf
	Here we ignore the number of bytes printed by assigning it to _ */
	_, err := fmt.Printf("%s", msg)
	return err
}

func main() {
	printer("Hello, World!")
}
