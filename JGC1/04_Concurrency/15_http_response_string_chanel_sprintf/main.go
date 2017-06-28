package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/* getter takes the url string and a string channel to which it will send the lenght of the url */
func getter(url string, lengthsChan chan string) {
	length, err := getPage(url)
	if err == nil {
		/* godoc fmt Sprintf
		send the lenght of the url as string to the lengthsChan string channel */
		lengthsChan <- fmt.Sprintf("%s is %d bytes long", url, length)
	}
}

func getPage(url string) (int, error) {
	/* godoc net/http Get
	   godoc net/http Response */
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	/* defer closes the Response Body before returning */
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

	/* Make string channel */
	lengthsChan := make(chan string)

	/* Ranging over slice of strings returns index and the string at that index
	   Here we are ignoring the index by assigning it to _ */
	for _, url := range urls {
		/* Start a getter goroutine for each url */
		go getter(url, lengthsChan)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-lengthsChan)
	}
}
