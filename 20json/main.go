package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"Course Name"`
	Price    int    `json:"Course Price"`
	Platform string
	Password string   `json:"-"` // "-" ---> not include the field when send data
	Tags     []string `json:"tags,omitempty"`  // "omitempty" ---> not include the field if that is nil
}

func main() {
	encodeJSON()

}

func encodeJSON() {
	myCourse := []course{
		{"React", 299, "Youtube", "12345", []string{"JS", "Web-dev"}},
		{"Angular", 399, "Youtube", "09876", nil},
	}

	json, err := json.MarshalIndent(myCourse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}
