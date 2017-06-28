package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func worker(urlChan chan string, lengthsChan chan string, id int) {
	for {
		url := <-urlChan
		length, err := getPage(url)
		if err == nil {
			/* godoc fmt Sprintf */
			lengthsChan <- fmt.Sprintf("%s is %d bytes long : processed by %d", url, length, id)
		} else {
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
	// godoc io/ioutil ReadAll
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func main() {
	urlChan := make(chan string)
	lengthsChan := make(chan string)

	/* start 10 worker goroutines passing in worker id along with urlChan and lengthsChan */
	for i := 0; i < 10; i++ {
		go worker(urlChan, lengthsChan, i)
	}

	/* slice of strings */
	urls := []string{"http://www.oreilly.com/", "http://www.google.com/", "http://www.yahoo.com", "http://www.bing.com", "http://bbc.co.uk"}

	/* ranging over slice of string returns index and value at that index
	here we are ignoring the index by assigning it to _ */
	for _, url := range urls {
		/* send each url in to urlChan channel */
		urlChan <- url
	}

	/* receive and print the results */
	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-lengthsChan)
	}
}
