package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const Url string = "https://cat-fact.herokuapp.com/facts"

const postUrl string = "https://httpbin.org/post"

func main() {
	getRequest(Url)
	postRequest(postUrl)

}

func getRequest(url string) {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var responseString strings.Builder
	data, _ := io.ReadAll(res.Body)
	responseString.Write(data)

	fmt.Println(responseString.String())
}

func postRequest(url string) {
	requestBody := strings.NewReader(`
		{
			"name": "Alice Johnson",
			"age": 28,
			"skills": ["Python", "JavaScript", "Django", "React"],
			"currentProject": "EcoTrack - A carbon footprint tracking app",
			"email": "alice.j@example.com",
			"location": "San Francisco, CA",
			"github": "https://github.com/alice_j"
		}
	`)
	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content,_ :=io.ReadAll(response.Body)

	fmt.Println(string(content))
}
