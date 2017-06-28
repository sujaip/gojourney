package main

/* godoc fmt Errorf */

import (
	"fmt"
)

func main() {
	var err error
	thing := "My Module"
	err = fmt.Errorf("A Problem occured with %s", thing)
	fmt.Printf("Error: %s\n", err)
}
