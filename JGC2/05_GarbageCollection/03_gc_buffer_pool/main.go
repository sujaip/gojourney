package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	/* make a byte slice of random size */
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	pool := make([][]byte, 200)

	for i := 0; i < 10; i++ {
		go func(offset int) {
			for {
				b := makeBuffer()
				j := offset + rand.Intn(20)
				pool[j] = b
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			}
		}(i * 20)
	}

	for {
		time.Sleep(1 * time.Second)
		var r runtime.MemStats
		runtime.ReadMemStats(&r)
		fmt.Printf("Heap Alloc = %d\n", r.HeapAlloc)
	}
}

/*
Heap allocation for this program grows and shrinks rapidly due to
garbage collection of unused buffers pool of buffers

Common way of keeping the heap size small is to recycle the pool of buffers
i.e. make the unused buffers available to some other goroutine for reuse rather than garbage collecting it
*/
