package main

import (
	"fmt"
	"time"
)

func makeID(idChan chan int) {
	id := 0

	for {
		idChan <- id
		id++
	}
}

func main() {
	ids := make(chan int)

	/* start two makeID goroutines
	   each makeID goroutine will send its own series of ids */
	go makeID(ids)
	go makeID(ids)

	/* <-ids here means to get next value from the ids channel */
	fmt.Printf("%d ", <-ids)

	time.Sleep(1 * time.Second)

	/* here we are using the ids channel to get next id from makeID goroutiene */
	for id := range ids {
		fmt.Printf("%d ", id)
		time.Sleep(100 * time.Millisecond)
	}

	/* Output:
	1 0 2 1 3 2 4 3 5 4 6 5 7 6 8 7 9 8 10 9 11 10 12 11 13 12 14 13 15 14 16 15 ...... */
}
