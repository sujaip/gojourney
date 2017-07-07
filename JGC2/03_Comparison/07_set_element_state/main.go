package main

import "fmt"

type Device struct {
	Name string
}

func main() {
	gadgets := make(map[Device]bool)
	gadgets[Device{"iphone"}] = true
	gadgets[Device{"imac"}] = false

	present, ok := gadgets[Device{"imac"}]
	if ok {
		fmt.Printf("%t\n", present)
	} else {
		fmt.Printf("Can't find the item")
	}
}

/*
Sometimes we might want to check if each element in the set is onn or off
*/
