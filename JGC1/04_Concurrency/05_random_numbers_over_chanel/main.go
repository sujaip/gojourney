package main

import (
	"fmt"
	"math/rand"
)

func makerandoms(c chan int) {
	for {
		/* godoc math/rand Intn */
		c <- rand.Intn(1000)
	}
}

func main() {
	randoms := make(chan int)

	go makerandoms(randoms)

	/* receive random numbers from the channel and keep printing forever */
	for n := range randoms {
		fmt.Printf("%d ", n)
	}
}
