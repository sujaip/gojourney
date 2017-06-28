package main

import "fmt"

func main() {
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	has28 := 0

	delete(dayMonths, "Feb")

	/* Deleting second time does not cause any error */
	delete(dayMonths, "Feb")

	/* Deleting elements that don't exist is fine as well */
	delete(dayMonths, "February")

	for _, days := range dayMonths {
		if days == 28 {
			has28++
		}
	}

	fmt.Printf("%d months have 28 days\n", has28)
}
