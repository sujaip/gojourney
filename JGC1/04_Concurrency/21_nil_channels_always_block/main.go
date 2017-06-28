package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* nil channel will block when we try to receive from it
   this behaviour is used to turn off a part of the select
	 statement by setting the channel to nil
	 http://www.godesignpatterns.com/2014/05/nil-channels-always-block.html */
func reader(ch chan int) {
	t := time.NewTimer(2 * time.Second)
	for {
		select {
		/* this case is ignored after the ch channel is set to nil */
		case i := <-ch:
			fmt.Printf("%d\n", i)
		case <-t.C:
			/* set ch channel to nil after three seconds */
			ch = nil
		}
	}
}

func writer(ch chan int) {
	for {
		/* godoc math/rand Intn */
		ch <- rand.Intn(42)
	}
}

func main() {
	ch := make(chan int)
	go reader(ch)
	go writer(ch)

	time.Sleep(3 * time.Second)
}
