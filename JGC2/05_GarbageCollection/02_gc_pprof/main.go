package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func save(sizes []int, n int) []int {
	return append(sizes, n)
}

func main() {
	s := "Hello, World!"
	var sizes []int

	first := true
	start := time.Now()

	for {
		if time.Since(start) > time.Second {
			var r runtime.MemStats
			runtime.ReadMemStats(&r)
			fmt.Printf("Heap size %d\n", r.HeapAlloc)
			fmt.Printf("NumGC %d\n", r.NumGC)
			start = time.Now()
			if first {
				out, _ := os.Create("profile")
				pprof.WriteHeapProfile(out)
				out.Close()
				first = false
			}
		}
		sizes = save(sizes, len(s))
	}
}

/*
$ go build main.go
$ ./main
Stop the program after a moment, it would have created the "profile" file
$ ls profile
profile

Below command will bringup the pprof prompt
$ go tool pprof main profile
...
(pprof) top
(pprof) list
(pprof) web
*/
