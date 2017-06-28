package main

import (
	"fmt"
	"time"
)

/* printer goroutine keeps printing till something is received on stopChan */
func printer(msg string, stopChan chan bool) {
	for {
		select {
		/* receives a zero value when the stopChan is closed */
		case <-stopChan:
			return
		default:
			fmt.Printf("%s\n", msg)
		}
	}
}

func main() {
	stopChan := make(chan bool)

	/* start ten printer goroutines that will keep printing till the stopChan channel is closed */
	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("printer: %d", i), stopChan)
	}

	time.Sleep(2 * time.Second)
	/* close is received by all the waiting goroutines
	close is used as a broadcast here to indicate every one should stop */
	close(stopChan)
}
