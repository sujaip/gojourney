package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func makeBuffer() interface{} {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	pool := make([][]byte, 200)
	var spare sync.Pool
	spare.New = makeBuffer

	for i := 0; i < 10; i++ {
		go func(offset int) {
			for {
				b := spare.Get().([]byte)
				j := offset + rand.Intn(20)
				if pool[j] != nil {
					spare.Put(pool[j])
				}
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
Sync Pool's purpose is to cache allocated but unused items for later reuse.
This relieves pressure on the Garbage Collector
Sync pool is safe for use by multiple goroutines

However any item stored in the pool might be removed automatically at any time without notification
If the only reference to the item is held by the pool when this happens then the item might be dealocated
*/
