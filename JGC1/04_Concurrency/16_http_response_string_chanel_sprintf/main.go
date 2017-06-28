package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func worker(urlChan chan string, lengthsChan chan string) {
	for {
		url := <-urlChan
		length, err := getPage(url)
		if err == nil {
			/* godoc fmt Sprintf */
			lengthsChan <- fmt.Sprintf("%s is %d bytes long", url, length)
		} else {
			/* send the error message on lengthsChan in case of error */
			lengthsChan <- fmt.Sprintf("Error getting %s: %s", url, err)
		}
	}
}

func getPage(url string) (int, error) {
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
	urlChan := make(chan string)
	lengthsChan := make(chan string)

	/* start the worker goroutine
	worker goroutine waits to receive url on the urlChan string channel */
	go worker(urlChan, lengthsChan)

	/* send url string to the urlChan string channel */
	urlChan <- "http://www.cluvd.io/"
	/* receive from lengthsChan and print the results */
	fmt.Printf("%s\n", <-lengthsChan)
}
