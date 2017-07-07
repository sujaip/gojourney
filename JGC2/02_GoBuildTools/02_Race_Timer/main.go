package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomDuration() time.Duration {
	/* godoc math/rand Int63n
	func Int63n(n init64) int64
	returns an int64 non-negative pseudo random number */
	return time.Duration(rand.Int63n(1e9))
}

func main() {
	start := time.Now()
	var t *time.Timer
	/* Call a function after random amount of time
	   called function resets the timer for some time in the future
	   godoc time AfterFunc */
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
