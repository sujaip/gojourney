package main

import "fmt"

func do() {
	defer fmt.Printf("First\n")
	defer fmt.Printf("Second\n")
	defer fmt.Printf("Third\n")
}

func main() {
	do()
}

/*
This program output is as given below:
Third
Second
First
Defer is Last in First Out
***************************/
