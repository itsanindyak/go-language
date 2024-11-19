package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"Course Name"`
	Price    int    `json:"Course Price"`
	Platform string
	Password string   `json:"-"`              // "-" ---> not include the field when send data
	Tags     []string `json:"tags,omitempty"` // "omitempty" ---> not include the field if that is nil
}

var jsonDataFromWeb = []byte(`
	{
        "Course Name": "React",
        "Course Price": 299,
        "Platform": "Youtube",
        "tags": ["JS","Web-dev"]
    }
`)

func main() {
	encodeJSON()
	decodeJSON(jsonDataFromWeb)
	
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

func decodeJSON(jsonData []byte)  {
	// var decodejson course

	// isValid := json.Valid(jsonData)

	// if isValid {
	// 	json.Unmarshal(jsonData, &decodejson)
	// }



	// add data  to key value

	var decodejson map[string]interface{}

	json.Unmarshal(jsonData,&decodejson)


	for key,value := range decodejson{

		fmt.Println(key,"=",value)
	}
}
