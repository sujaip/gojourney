package main

import "fmt"

func main() {
	/* Map Literal */
	dayMonths := map[string]int{"Jan": 31, "Feb": 28, "Mar": 31, "Apr": 30, "May": 31,
		"Jun": 30, "Jul": 31, "Aug": 31, "Sep": 30, "Oct": 31, "Nov": 30, "Dec": 31,
	}

	has28 := 0

	for _, days := range dayMonths {
		if days == 28 {
			has28++
		}
	}

	fmt.Printf("%d months have 28 days\n", has28)
}
