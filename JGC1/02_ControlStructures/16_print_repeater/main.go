package main

import "fmt"

func printer(msg1, msg2 string, repeate int) {
	for repeate > 0 {
		fmt.Printf("%s", msg1)
		fmt.Printf("%s\n", msg2)
		repeate--
	}
}

func main() {
	printer("Hello,", " World!", 4)
}
