package main

import (
	"fmt"
)

func main() {
	b := byte(65)
	/* %x here in the Printf means to print it as hex value (base 16)*/
	fmt.Printf("Value: %x\n", b)
}
