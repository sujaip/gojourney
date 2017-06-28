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

	/* delete removes the coresponding key and value from the map */
	delete(dayMonths, "Feb")

	has28 := 0
	/* In thei case we have ignored the keys by assigning to _
	   if not Go will complain about the months variable not beeing used */
	for _, days := range dayMonths {
		if days == 28 {
			has28++
		}
	}

	fmt.Printf("%d months have 28 days\n", has28)
}
