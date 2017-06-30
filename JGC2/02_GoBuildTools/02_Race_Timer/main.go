package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

func main() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

/* occationaly this program might crash due to race condition
   we can use the -race flag with go run to help detect issues with race
   $ go run -race main.go
*/
