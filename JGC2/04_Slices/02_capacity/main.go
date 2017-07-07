package main

import "fmt"

func main() {
	mySlice := []int{2, 4, 6, 10, 12, 14}
	sliceOfmySlice := mySlice[1:3]
	fmt.Printf("%v\n", sliceOfmySlice)

	mySlice[1] *= 2
	fmt.Printf("%v\n", sliceOfmySlice)
	/*
		Go does not have any inbuilt mechanism to remove an element from middle of a slice
	*/

	sliceOfmySlice = append(sliceOfmySlice, 8)
	fmt.Printf("%v\n", sliceOfmySlice)
	fmt.Printf("%v\n", mySlice)

	/*
		Appending a slice beyond the capacity will create a new copy of the slice
		Once the slice of slice is expanded beyond its capacity,
		it becomes different from the underlying slice and operations on it don't affect the underlying slice
	*/
	fmt.Printf("Capacity of sliceOfmySlice: %d\n", cap(sliceOfmySlice))
	sliceOfmySlice = append(sliceOfmySlice, 16)
	sliceOfmySlice = append(sliceOfmySlice, 18)
	sliceOfmySlice = append(sliceOfmySlice, 20)
	sliceOfmySlice = append(sliceOfmySlice, 22)
	fmt.Printf("%v\n", sliceOfmySlice)
	fmt.Printf("%v\n", mySlice)
	fmt.Printf("Capacity of sliceOfmySlice: %d\n", cap(sliceOfmySlice))

}
