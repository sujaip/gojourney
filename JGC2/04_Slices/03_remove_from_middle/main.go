package main

import "fmt"

func main() {
	mySlice := []int{2, 4, 8, 6, 10, 12, 14}

	/*
		Go does not have any inbuilt mechanism to remove an element from middle of a slice
	*/

	fmt.Printf("%v\n", mySlice)
	mySlice = append(mySlice[:3], mySlice[4:]...)
	fmt.Printf("%v\n", mySlice)

	/*
		mySlice[:3] -> From begining of the slice upto but not including 3rd element
		mySlice[4:] -> From 4th element of the slice till last element
	*/

	/*
		Three forms of append
		(1.) slice = append(slice, element1, element2)
		(2.) slice = append(slice, anotherSlice)
		(3.) slice = append([]byte("hello"), "world")
		Third form is to append a string to a byte slice
	*/

}
