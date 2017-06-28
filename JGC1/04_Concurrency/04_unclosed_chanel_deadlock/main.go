package main

import (
	"fmt"
	"time"
)

func emit(c chan string) {
	words := []string{"The", "Quick", "Brown", "Fox"}
	for _, word := range words {
		c <- word
		time.Sleep(2 * time.Millisecond)
	}
	/*close(c)
	fatal error: all goroutines are asleep will not occur if this channel is closed after exiting the for loop
	but we don't want one of the emit goroutines to close the channel before the other one is finished */
}

func main() {
	wordChannel := make(chan string)

	/* start two emit goroutines */
	go emit(wordChannel)
	go emit(wordChannel)

	/* fatal error: all goroutines are asleep - deadlock!
	This error occurs as range is going to wait for ever trying to read from the empty unclosed channel
	"No one to close the channel" */
	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}

	/* Output:
	The The Quick Quick Brown Fox Brown Fox fatal error: all goroutines are asleep - deadlock! */
}
