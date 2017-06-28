package main

import "fmt"

type Translation struct {
	English string
	French  string
}

func main() {
	title1 := Translation{"Mister", "Monsieur"}
	/* Anonymous Struct */
	title2 := struct {
		English string
		French  string
	}{"Mister", "Monsieur"}

	fmt.Printf("%#v\n", title1)
	fmt.Printf("%#v\n", title2)

}
