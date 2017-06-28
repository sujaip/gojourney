package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Buffered channel Example */
func work() {
	fmt.Printf("[")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("]")
}

func worker(sema chan bool, id int) {
	<-sema
	work()
	fmt.Printf("Worker %d\n", id)
	sema <- true
}

func main() {
	/* making buffered channel */
	sema := make(chan bool, 100)

	/* start thousand worker goroutines */
	for i := 0; i < 1000; i++ {
		go worker(sema, i)
	}

	/* cap(sema) gives the capacity of the sema channel */
	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	time.Sleep(30 * time.Second)
}
