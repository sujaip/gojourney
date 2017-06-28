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

/* pointer receiver */
func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()
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
	//w := &webPage{url: "http://www.cluvd.io/"}
	//w := new(webPage)
	//w := &webPage{}
	w := webPage{}
	w.url = "http://www.cluvd.io/"
	w.get()
	if w.isOK() {
		fmt.Printf("URL: %s Length: %d\n", w.url, len(w.body))
	} else {
		fmt.Printf("Something went wrong: %s\n", w.err)
	}
}
