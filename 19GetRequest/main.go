package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const Url string = "https://cat-fact.herokuapp.com/facts"

func main() {
	getRequest(Url)

}

func getRequest(url string) {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	var responseString strings.Builder
	data, _ := io.ReadAll(res.Body)
	responseString.Write(data)

	fmt.Println(responseString.String())

}
