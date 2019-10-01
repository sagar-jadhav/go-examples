package main

import (
	"fmt"
	"net/http"
)

type response struct {
	Code int
	URL  string
	Err  error
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://archlinux.org",
		"invalidurl",
	}

	responses := make(chan response)

	for _, url := range urls {
		go checkURL(url, responses)
	}

	for range urls {
		r := <-responses
		if r.Err != nil {
			fmt.Printf("Got error when checking %s: %s\n", r.URL, r.Err)
		} else {
			fmt.Printf("%s:\t%d\n", r.URL, r.Code)
		}
	}
}

func checkURL(url string, responses chan response) {
	resp, err := http.Get(url)
	if err != nil {
		responses <- response{
			URL: url,
			Err: err,
		}
		return
	}

	responses <- response{
		URL:  url,
		Code: resp.StatusCode,
	}
}
