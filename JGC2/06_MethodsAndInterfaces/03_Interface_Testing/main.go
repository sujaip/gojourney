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
		Here we are assigning ls to an anonymous variable of Yeller Interface
	*/
	var _ yelling.Yeller = ls
}

/*
This program gives out the following error when trying to build it
        cannot use ls (type *yelling.LoudString) as type yelling.Yeller in assignment:
        *yelling.LoudString does not implement yelling.Yeller (missing Len method)

We have to implement the Len method on LoufString in yelling package to fix this error

    func (ls *LoudString) Len() int {
      return len(ls.s)
    }

*/
