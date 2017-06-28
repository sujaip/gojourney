package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var running int64

func work() {
	/* godoc sync/atomic AddInt64
	increment running by one, notice the pass by reference */
	atomic.AddInt64(&running, 1)
	/* Prints the number of currently running goroutines */
	fmt.Printf("> %d ", running)

	/* godoc time Duration
	Here we are converting random int to time.Duration */
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	/* If we don't mimic work with the above time.Sleep, then we will always see 2
	or 3 as the ouput as goroutines exit immediately after decementing the running count */

	fmt.Printf("<")
	/* Decrement running using negative addition */
	atomic.AddInt64(&running, -1)
}

func worker(sema chan bool) {
	<-sema
	work()
	sema <- true
}

func main() {
	/* making buffered channel */
	sema := make(chan bool, 100)

	for i := 0; i < 1000; i++ {
		go worker(sema)
	}

	/* send true over buffered channel sema */
	for i := 0; i < 20; i++ {
		sema <- true
	}

	time.Sleep(5 * time.Second)
}
