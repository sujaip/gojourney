package main

import (
	"errors"
	"fmt"
)

func do() {
	var err error
	defer func() {
		if err != nil {
			fmt.Printf("An error occured: %s\n", err)
		}
	}()
	err = errors.New("Some Badass Error")
}

func main() {
	do()
}

/*
This program output is as given below:
An error occured: Some Badass Error

The actual function called by defer is evaluated when the defer is executed
***************************************************************************/
