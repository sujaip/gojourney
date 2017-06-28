package main

/* godoc fmt Errorf */

import (
	"fmt"
	"time"
)

type Error interface {
	Error() string
}

type MyError struct {
	str  string
	when time.Time
}

func (m MyError) Error() string {
	return fmt.Sprintf("%s at %s\n", m.str, m.when)
}

func main() {
	err := MyError{str: "Bad thing", when: time.Now()}
	fmt.Printf("%s", err)
}

/* Error itself is actually just an Interface with a single Error method that returns a string
Any type that has a Error function is an error*/
