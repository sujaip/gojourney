package main

import (
	"fmt"
	"time"
)

func makeID(idChan chan int) {
	var id int
	id = 0

	for {
		idChan <- id
		id++
	}
}

func main() {
	ids := make(chan int)

	go makeID(ids)

	for {
		/* We can use the receive from channel operation inside any compound statement */
		fmt.Printf("%d ", <-ids)
		time.Sleep(10 * time.Millisecond)
	}
}
