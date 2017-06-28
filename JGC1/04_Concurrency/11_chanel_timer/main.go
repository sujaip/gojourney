package main

import (
	"fmt"
	"time"
)

func emit(wordChannel chan string) {
	/* defer will close the wordChannel before returning */
	defer close(wordChannel)
	words := []string{"The", "quick", "brown", "fox"}
	i := 0

	/* godoc time Timer
	   Timer sends on t.C channel after timeout
		 When the Timer expires, current time will be sent on C */
	t := time.NewTimer(5 * time.Second)

	for {
		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		/* When the Timer expires, the current time will be sent on C
		and the emit goroutine will return */
		case <-t.C:
			return
		}
	}
}

func main() {
	wordCh := make(chan string)

	go emit(wordCh)

	/* range over the wordCh channel till the timer expires */
	for word := range wordCh {
		fmt.Printf("%s ", word)
	}

}
