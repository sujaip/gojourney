package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	/*Byte Slice of length 100 */
	b := make([]byte, 100)
	fmt.Printf("\nLength of byte slice b %d\n", len(b)) // Length of byte slice b 100

	/* when we do a := with err like below
	go will reuse the previously declared err variable */
	n, err := f.Read(b) // here n holds the number of bytes read

	/* print raw bytes */
	fmt.Printf("%d: % x\n", n, b)

	/* print string version of the bytes */
	fmt.Printf("%d: %s\n", n, string(b))

	/* string(b) here converts the byte slice in to string
	[]byte(someString) will conver a string in to byte slice */

	/* cannot convert b (type []byte) to type int
	fmt.Printf("%d: %\n", n, int(b)) */
}
