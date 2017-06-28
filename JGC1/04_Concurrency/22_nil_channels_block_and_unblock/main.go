package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reader(ch chan int) {
	t := time.NewTimer(8 * time.Second)
	for {
		select {
		case i := <-ch:
			fmt.Printf("%d\n", i)
		case <-t.C:
			/* Setting channel to nil after 8 seconds */
			ch = nil
		}
	}
}

func writer(ch chan int) {
	/* Setting channel to nil after 2 seconds and turning it on again after 6 seconds */
	stopper := time.NewTimer(2 * time.Second)
	restarter := time.NewTimer(6 * time.Second)

	/* Save the channel for future */
	savedCh := ch

	for {
		select {
		case ch <- rand.Intn(42):
		case <-stopper.C:
			ch = nil
		case <-restarter.C:
			ch = savedCh
		}
	}
}

func main() {
	ch := make(chan int)
	go reader(ch)
	go writer(ch)
	time.Sleep(10 * time.Second)
}
