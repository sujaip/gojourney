package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/* User-Defined Types */

type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)
	defer resp.Body.Close()
	if err != nil {
		w.err = err
		return
	}
	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
		return
	}
}

func (w *webPage) isOK() bool {
	return w.err == nil
}

func main() {
	w1 := &webPage{}
	w2 := new(webPage)
	w1.url = "http://www.oreilly.com/"
	w2.url = "http://www.google.com/"
	w1.get()
	if w1.isOK() {
		fmt.Printf("URL: %s Error: %s Length: %d\n", w1.url, w1.err, len(w1.body))
	} else {
		fmt.Printf("Something went wrong\n")
	}
	w2.get()
	if w2.isOK() {
		fmt.Printf("URL: %s Error: %s Length: %d\n", w2.url, w2.err, len(w2.body))
	} else {
		fmt.Printf("Something went wrong\n")
	}
}
