package main

import (
	"fmt"
	"yelling"
)

func main() {
	fmt.Printf("Copy Receiver or Pass By Value: ")
	lsv := yelling.NewLoudStirng()
	lsv.ValueChange("Hello")
	fmt.Println(lsv.String())

	fmt.Printf("Pointer Receiver or Pass By Reference: ")
	lsp := yelling.NewLoudStirng()
	lsp.PointerChange("Hello")
	fmt.Println(lsp.String())
}
