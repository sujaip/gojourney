package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/* takes url as string parameter and returns the length as int and error status */
func getPageLength(url string) (int, error) {
	/* godoc net/http Get
	   godoc net/http Response */
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	/* defer will close the response body before returning */
	defer resp.Body.Close()

	/* godoc io/ioutil ReadAll */
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func main() {
	/* slice of strings */
	urls := []string{"http://www.google.com/", "http://www.yahoo.com", "http://www.bing.com", "http://bbc.co.uk"}

	for _, url := range urls {
		pageLength, err := getPageLength(url)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("%s is %d bytes long\n", url, pageLength)
	}
}
