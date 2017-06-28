package main

import (
	"fmt"
	"time"
)

func emit(chanOfChan chan chan string) {
	/* Make a sting channel */
	wordChan := make(chan string)
	/* send the wordChan in to channel of string channels */
	chanOfChan <- wordChan
	/* defer will close the string channel before returning */
	defer close(wordChan)
	/* slice of strings */
	words := []string{"The", "quick", "brown", "fox"}

	i := 0
	/* godoc time Timer
	   This timer sends on the t.C channel after the timeout */
	t := time.NewTimer(1 * time.Second)

	for {
		select {
		case wordChan <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-t.C:
			/* When the Timer expires, the current time will be sent on C */
			return
		}
	}
}

func main() {
	/* Make channel of string channels */
	chanOfChan := make(chan chan string)

	/* start the emit goroutine and pass channel of string channels as argument */
	go emit(chanOfChan)

	/* receive the words string chanel from channel of string channels */
	words := <-chanOfChan

	/* range over string channel */
	for word := range words {
		fmt.Printf("%s ", word)
	}

}
