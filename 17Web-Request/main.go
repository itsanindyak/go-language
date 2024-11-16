package main

import (
	"fmt"
	"io"
	"net/http"
)

const Url string = "https://cat-fact.herokuapp.com/facts"

func main() {
	response, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
