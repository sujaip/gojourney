package main

import (
	"fmt"
)

func main() {
	var pi float64 = 3.14
	/* we could also declare it with short assingment like below
	pi := float64(3.14)
	here the float64() explicitly mentions the type*/

	fmt.Printf("Value: %f\n", pi)
	/* we can use the %.2f format string to print only two decimal points*/
	fmt.Printf("Value: %.2f\n", pi)
}
