package main

import "fmt"

func emit(wordChannel chan string, done chan bool) {
	words := []string{"The", "quick", "brown", "fox"}
	i := 0
	for {
		select {
		/* send on wordsCh chanel */
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-done:
			/* done <- true
			   Run time error: fatal error: all goroutines are asleep - deadlock! */
			return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for i := 0; i < 8; i++ {
		fmt.Printf("%s ", <-wordCh)
	}

	/* send on doneCh and wait for reply */
	doneCh <- true
	<-doneCh
	/* Run time error: fatal error: all goroutines are asleep - deadlock!
	   As we are waiting to recive on doneCh and no one is sending on it */

}
