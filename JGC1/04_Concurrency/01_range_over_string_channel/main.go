package main

import "fmt"

func emit(c chan string) {
	/*Slice of strings */
	words := []string{"The", "Quick", "Brown", "Fox"}

	/*for each string in the slice of strings */
	for _, word := range words {
		c <- word
		/* c <- here means send it in to the channel */
	}
	/* close the channel */
	close(c)
}

func main() {
	/* make a channel of type string */
	wordChannel := make(chan string)

	/* start emit as a goroutiene */
	go emit(wordChannel)

	/* range over the channel and print the strings as they arive */
	for word := range wordChannel {
		fmt.Printf("%s\n", word)
	}
}
