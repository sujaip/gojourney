package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	/* godoc net/http Get
	   http.Get takes a url string
	   http.Get returns a Pointer to Response struct and an error
	   https://golang.org/pkg/net/http/#Get
	   https://golang.org/pkg/net/http/#Response */
	resp, _ := http.Get("http://example.com/")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
