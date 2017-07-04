package main

import "fmt"

/* Named return values first and second */
func Names() (first, second string) {
	first = "Foo"
	second = "Bar"
	return
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)
}
