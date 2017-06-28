package main

/* play.golang.org */
import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, World!\n")
	/* Upper case P in the Printf indicates that it is exported */

	/* Within Go strings are based on runes which are unicode characters
	Its even possible to use Unicode characters in function names*/
}

/* $ export GOPATH=`pwd`
$ echo $GOPATH
/home/vagrant/gohome
$ pwd
/home/vagrant/gohome */

/* $ go fmt hello */
