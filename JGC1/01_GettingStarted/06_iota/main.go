package main

import (
	"fmt"
)

/* Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants.
   It is reset to 0 whenever the reserved word const appears in the source */

const (
	message  = "%d %d\n"
	messages = "%d %d\n"
	answer1  = iota
	answer2
)

const (
	second_message = "%d %d %d\n"
	second_answer1 = iota * 2
	second_answer2
	second_answer3
)

func main() {
	fmt.Printf(message, answer1, answer2)
	fmt.Printf(messages, answer1, answer2)
	fmt.Printf(second_message, second_answer1, second_answer2, second_answer3)
}

/* Output of this program would be like below
2 3
2 3
2 4 6

iota starts from 0, but when there are other assignments before it
then it starts from that position
*/
