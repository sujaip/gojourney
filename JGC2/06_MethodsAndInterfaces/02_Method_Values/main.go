package main

import (
	"fmt"
	"yelling"
)

func main() {
	/*
		Go allows us to use methods as values like listed below
	*/
	nls := yelling.NewLoudStirng
	ls := nls()

	change := ls.PointerChange
	change("Be Quiet")
	fmt.Println(ls.String())

}
