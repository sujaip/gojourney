package main

import (
	"fmt"
	"yelling"
)

/*

There are scenarios where we want to use a sub set of an interface methods without completely
implementing all the methods of the interface on the concrete type

We can explicitly declare like below
This is mostly used in test cases where we want to test few methods and not the full interface

In this example we have removed the Len method from yelling package
We have also explicitly declared LoudString to be of Yeller type like below

type LoudString struct {
	Yeller
	s string
}

*/

func main() {
	ls := yelling.NewLoudStirng()
	ls.Change("Hello")
	fmt.Println(ls.String())

	var _ yelling.Yeller = ls

	// fmt.Printf("%d\n", ls.Len())
	// We will get the below run time error if we uncomment this line
	// invalid memory address or nil pointer dereference
}
