package main

import "fmt"

func printer(msg1, msg2 string) {
	fmt.Printf("%s", msg1)
	fmt.Printf("%s\n", msg2)
}

func main() {
	printer("Hello,", " World!")
}
