package main

import (
	"fmt"
	"time"
)

func printer(msg string, goChan chan bool) {
	/* this receives a zero value when the channel is closed */
	<-goChan
	fmt.Printf("%s\n", msg)
}

func main() {
	goChan := make(chan bool)

	/* start ten printer goroutines that will wait on goChan channel */
	for i := 0; i < 10; i++ {
		/* send loop index as printer id for each printer goroutine */
		go printer(fmt.Sprintf("printer: %d", i), goChan)
	}

	time.Sleep(1 * time.Second)
	/* close is received by all the waiting goroutines */
	close(goChan)
	time.Sleep(1 * time.Second)

}
