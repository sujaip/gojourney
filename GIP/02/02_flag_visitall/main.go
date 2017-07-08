package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
}

func main() {
	flag.Parse()

	/* godoc flag VisitAll
	   flag package has a VisitAll function that accepts a callback function as an argument.
	   VisitAll iterates over each of the flags executing the callback function on it
	   We can use VisitAll function to write out our own help text for each flag*/
	flag.VisitAll(func(flag *flag.Flag) {
		format := "\t-%s: %s (Default: '%s')\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
	})
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}

/*
https://godoc.org/launchpad.net/gnuflag
Above gnuflag package has almost exact same API as standard flag package
With one additional first argument to Parse function of either true or false
A value of true looks for flags anywhere in the command, and false behaves like standard flag package

https://github.com/jessevdk/go-flags
Above go-flags package provides Linux and BSD style flags, but uses entirely different API and style
*/
