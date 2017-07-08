package main

import "fmt"

type Counter struct {
	c    chan int
	done chan struct{}
	i    int
}

func NewCounter() *Counter {
	counter := new(Counter)
	counter.c = make(chan int)
	counter.done = make(chan struct{})

	/*
		Start a goroutine that keeps sending integers on channel c until it receives from the done channel
	*/
	go func() {
		for {
			select {
			case counter.c <- counter.i:
				counter.i++
			case <-counter.done:
				return
			}
		}
	}()

	return counter
}

func (c *Counter) GetSource() <-chan int {
	return c.c
}

func (c *Counter) Stop() {
	c.done <- struct{}{}
}

func main() {
	c := NewCounter()
	read := c.GetSource()
	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	c.Stop()
}
