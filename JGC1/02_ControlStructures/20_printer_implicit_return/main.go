package main

import "fmt"

/* since the error return value is named e,
we don't have to explicitly return it */
func printer(msg string) (e error) {
	_, e = fmt.Printf("%s\n", msg)
	return
}

func main() {
	err := printer("Hello, World!")
	if err == nil {
		fmt.Printf("Success")
	}
}
