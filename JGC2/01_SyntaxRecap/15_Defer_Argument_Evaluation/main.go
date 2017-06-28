package main

import "fmt"

func do() {
	n := 1
	defer fmt.Printf("First (%d)\n", n)
	n = 2
	defer fmt.Printf("Second (%d)\n", n)
	n = 3
	defer fmt.Printf("Third (%d)\n", n)
}

func main() {
	do()
}

/*
This program output is as given below:
Third (3)
Second (2)
First (1)

Right hand side of defer is almost always a function call

Agrguments to functions called by defer are evaluated when the defer statement apears
The actual function called by defer is evaluated when the defer is executed
https://blog.golang.org/defer-panic-and-recover
*************************************************************************************/
