package main

import "fmt"

func emit(wordChannel chan string, done chan bool) {
	words := []string{"The", "quick", "brown", "fox"}
	i := 0
	for {
		select {
		/* Keep sending words on the wordChannel */
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		/* Waiting to receive on done chanel */
		case <-done:
			fmt.Printf("\nGot Done!\n")
			close(done)
			return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	/* Receive and print from word chanel */
	for i := 0; i < 16; i++ {
		fmt.Printf("%s ", <-wordCh)
	}

	/* Sending to doneCh to signal that we are done */
	doneCh <- true

}
