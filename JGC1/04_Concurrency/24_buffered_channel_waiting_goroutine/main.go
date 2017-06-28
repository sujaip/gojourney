package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Buffered channel Example */
func work(i int) {
	fmt.Printf("> %d ", i)
	/* Sleep for random seconds of time
	   Intn returns as an int a non-negative pseudo-random number in between 0 and n
	   Intn It panics if n <= 0 */
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	fmt.Printf("<\n")
}

func worker(sema chan bool, i int) {
	/* worker blocks on sema channel here */
	<-sema
	work(i)
	sema <- true
}

func main() {
	/* make a buffered channel */
	sema := make(chan bool, 100)

	/* start 1000 workers */
	for i := 0; i < 1000; i++ {
		go worker(sema, i)
	}

	/* worker goroutienes are blocked on sema channel during this sleeping period */
	time.Sleep(1 * time.Second)

	/* send true on buffered channel sema
	   worker goroutienes are blocked till we send true over sema */
	for i := 0; i < 100; i++ {
		sema <- true
	}

	/* sleep for 30 seconds */
	time.Sleep(5 * time.Second)
}
