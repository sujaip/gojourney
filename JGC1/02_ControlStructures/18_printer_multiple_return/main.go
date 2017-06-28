package main

import "fmt"

func printer(msg string) (string, error) {
	msg += "\n"
	_, err := fmt.Printf("%s", msg)
	return msg, err
}

func main() {
	message, err := printer("Hello, World!")
	if err == nil {
		/* %q prints out a double-quoted string safely escaped with Go syntax */
		fmt.Printf("%q\n", message)

		/* %x prints out the string in hex (base 16) */
		fmt.Printf("%x\n", message)

		/* % x prints out the string in hex with space between each character */
		fmt.Printf("% x\n", message)
	}
}
