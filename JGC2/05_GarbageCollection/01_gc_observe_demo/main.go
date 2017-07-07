package main

import (
	"fmt"
	"runtime"
	"time"
)

func length(s string) int {
	c := []byte(s)
	n := len(c)
	b := make([]byte, n)
	/* Above line creates garbage */
	return len(b)
}

/*
go run -gcflags="-m" main.go
look for lines with "escapes to heap"
*/

func main() {
	s := "Hello, World!"
	s += s
	s += s
	s += s

	start := time.Now()
	for {
		if time.Since(start) > time.Second {
			var r runtime.MemStats
			runtime.ReadMemStats(&r)
			fmt.Printf("Heap size %d\n", r.HeapAlloc)
			start = time.Now()
		}
		length(s)
	}
	/*
		This program generates garbage on each call to length,
		as a result the heap size is continously growing and shrinking
	*/
}
