package main

import "fmt"

func emit(c chan string) {
	/* slice of strings */
	words := []string{"The", "Quick"}
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

	/* receiving from channel also supports the (, ok) syntax for error checking */
	word, ok := <-wordChannel
	fmt.Printf("%s : %t\n", word, ok)

	/* receiving from closed channel we get zero value of the channel type */
	word, ok = <-wordChannel
	/* this will print empty string and false */
	fmt.Printf("%s : %t\n", word, ok)
}
