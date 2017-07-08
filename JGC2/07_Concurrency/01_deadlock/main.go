package main

import "fmt"

func main() {
	c := make(chan int)
	/*
		Here we are using an anonymous empty struct channel to signal work being done
		i.e. this channel is not going to be used for sending data, its sole puropse is to signal completion of work
	*/
	done := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case c <- i:
				i++
			case <-done:
				return
			}
		}
	}()

	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	/*
		Here we are sending an anonymous empty struct in to the done channel to indicate work is done
	*/
	done <- struct{}{}

	/*
		Below statement leads to RUN TIME error "all goroutines are asleep - deadlock!"
		This is due to the reason that we have already signaled the go routine to exit
		there is NO ONE TO SEND DATA ON CHANNEL c and we are going to be waiting forever - deadlock
	*/
	fmt.Printf("%d\n", <-c)
}
