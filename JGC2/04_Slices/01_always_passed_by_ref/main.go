package main

import "fmt"

func main() {
	mySlice := []int{2, 4, 6, 10, 12, 14}
	sliceOfmySlice := mySlice[1:3]
	fmt.Printf("%v\n", sliceOfmySlice)

	mySlice[1] *= 2
	fmt.Printf("%v\n", sliceOfmySlice)

	sliceOfmySlice = append(sliceOfmySlice, 8)
	fmt.Printf("%v\n", sliceOfmySlice)

	fmt.Printf("%v\n", mySlice)

	/*
		Changes to the underlying slice changes the copy as well
	*/

}

/*
Slices in go are pointer to an underlying array
Slices are always passed by reference in Go
Maps and channels are also always passed by reference in Go
*/
