package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wr sync.WaitGroup

func main() {
	websiteList := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com",
		"https://go.dev",
	}

	for _, website := range websiteList {
		wr.Add(1)
		go get(website)
	}

	wr.Wait()

}

func get(endpoint string) {
	defer wr.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error in ", endpoint)

	} else {
		fmt.Println(res.StatusCode, "for", endpoint)
	}
}
