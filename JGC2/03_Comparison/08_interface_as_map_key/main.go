package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var output = make(map[io.Writer]bool)

func main() {
	output[os.Stdout] = false
	output[os.Stderr] = true

	/*
	   godoc bytes Buffer
	   godoc bytes Write
	*/
	b := new(bytes.Buffer)
	output[b] = true

	for w, enabled := range output {
		if enabled {
			fmt.Fprintf(w, "Some Output\n")
		}
	}
	fmt.Printf("%d bytes written to Buffer\n", b.Len())

}

/*
Interfaces are comparable
Comparables can be used as map keys
*/
