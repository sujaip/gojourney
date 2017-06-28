package main

import "fmt"

func emit(c chan string) {
	/* slice of strings */
	words := []string{"The", "Quick", "Brown", "Fox"}
	for _, word := range words {
		c <- word
	}
	close(c)
}

func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)

	/* explicitly receive from the channel */
	word := <-wordChannel
	fmt.Printf("%s\n", word)

	word = <-wordChannel
	fmt.Printf("%s\n", word)

	/*Here we are exiting the main without receiving the third and fourth strings from the channle
	emit goroutine will get killed once we reach the end of main */

}
