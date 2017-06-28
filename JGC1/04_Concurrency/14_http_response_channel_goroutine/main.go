package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getter(url string, lengthsChan chan int) {
	length, err := getPageLength(url)
	if err == nil {
		lengthsChan <- length
	}
}

func getPageLength(url string) (int, error) {
	/* godoc net/http Get
	   godoc net/http Response */
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	/* godoc io/ioutil ReadAll */
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func main() {
	/* Slice of strings */
	urls := []string{"http://www.google.com/",
		"http://www.yahoo.com",
		"http://www.bing.com",
		"http://bbc.co.uk"}

	lengthsChan := make(chan int)

	/* Ranging over slice of strings returns index and the string at that index
	   Here we are ignoring the index by assigning it to _ */
	for _, url := range urls {
		/* start a getter goroutine for each url */
		go getter(url, lengthsChan)
	}

	/* Print lengths received on the channel */
	for i := 0; i < len(urls); i++ {
		fmt.Printf("Length %d\n", <-lengthsChan)
	}
}
