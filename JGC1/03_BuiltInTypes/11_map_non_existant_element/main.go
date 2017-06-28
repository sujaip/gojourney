package main

import "fmt"

func main() {
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31

	fmt.Printf("Days in February: %d\n", dayMonths["Februar"]) //Returns 0

	/* Accessing an element of a map that does not exist returns zero value of the type
	   We can safely access a non existant element without getting an error and rely on zero value */

	days, ok := dayMonths["January"]
	if !ok {
		fmt.Printf("Can't get days for January\n")
	} else {
		fmt.Printf("January has %d days\n", days)
	} /* Can't get days for January */

	days, ok = dayMonths["Jan"]
	if !ok {
		fmt.Printf("Can't get days for January\n")
	} else {
		fmt.Printf("January has %d days\n", days)
	} /* January has 31 days */

}
