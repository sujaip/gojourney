package main

import (
	"fmt"
	"time"
)

func emit(wordsCh chan string, done chan bool) {
	words := []string{"The", "quick", "brown", "fox"}
	i := 0
	for {
		select {
		/* send on wordsCh chanel */
		case wordsCh <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		/* receive on done chanel and send back on done channel */
		case <-done:
			time.Sleep(5 * time.Second)
			done <- true
			return
			/* goroutines can return as well */
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for i := 0; i < 16; i++ {
		fmt.Printf("%s ", <-wordCh)
	}

	/* send on doneCh and wait for reply */
	doneCh <- true
	<-doneCh

}
