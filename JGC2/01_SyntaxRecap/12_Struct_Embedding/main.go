package main

/***********************************************************************
One of the main uses of struct embedding is to embedd sync.Mutex in our
type so as to be able to call Lock and Unlock on instances of our type
*************************************************************************/

import (
	"fmt"
	"strings"
)

type MyString struct {
	str string
}

func NewMyString(s string) MyString {
	return MyString{str: s}
}

func (m MyString) Output() {
	fmt.Println(m.str)
}

type MyUpperString struct {
	MyString
}

func NewMyUpperString(s string) MyUpperString {
	mus := MyUpperString{}
	mus.str = strings.ToUpper(s)
	return mus
}

/* uncommenting this will override the Output method from embedded struct
func (m MyUpperString) Output() {
	fmt.Printf("Really Upper: %s\n", m.str)
} */

func main() {
	hello := NewMyString("Hello, World!")
	hello.Output()

	hi := NewMyUpperString("Hello, World!")
	/* Here Output is method from the embeded MyString struct */
	hi.Output()

}
