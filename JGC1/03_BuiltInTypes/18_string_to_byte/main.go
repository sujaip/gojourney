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

	//Byte Slice of 100
	b := make([]byte, 100)

	// when we do a := with err like below
	// go will reuse the previously declared err variable
	n, err := f.Read(b)

	fmt.Printf("%d: % x\n", n, b)

	stringVersionOfb := string(b)
	fmt.Printf("%d: %s\n", n, stringVersionOfb)

	// String can be converted back in to byte slice like below
	backToByteSlice := []byte(stringVersionOfb)
	fmt.Printf("%d: % x\n", n, backToByteSlice)
	//Type convertion between strings and byte slice is common in go

}
