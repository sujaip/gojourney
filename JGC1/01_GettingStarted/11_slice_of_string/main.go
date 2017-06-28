package main

import (
	"fmt"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	/* here atoz[0:9] means 9 characters from begining of the string
	we can also drop the 0 like atoz[:9] as default is from the begining of the string*/
	fmt.Printf("%s\n", atoz[0:9])

	/* atoz[16:19] means from 17th character till 19th character
	indexing begins from 0, so 16th character is excluded */
	fmt.Printf("%s\n", atoz[16:19])

	/* atoz[16:] means from 17th character till the end of the string */
	fmt.Printf("%s\n", atoz[16:])

}
