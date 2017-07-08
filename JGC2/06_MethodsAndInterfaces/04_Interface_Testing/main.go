package main

import (
	"fmt"
	"yelling"
)

func main() {
	ls := yelling.NewLoudStirng()
	ls.Change("Hello")
	fmt.Println(ls.String())

	/*
		Here we are assigning a value of type LoudString to an anonymous variable of Yeller Interface
	*/
	var _ yelling.Yeller = ls

	var _ fmt.Stringer = ls

	/*
		godoc fmt Stringer
		Stringer interface has one String() method that returns a string

		Here we are able to asssign value of type LoudString to an anonymous variable of Stringer interface
		As they both implement the single String() method from Stringer interface

	*/

}
