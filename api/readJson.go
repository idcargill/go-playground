package api

import (
	"encoding/json"
	"fmt"
	_ "fmt"
)

type Details struct {
	id int
	specs string
}

type Person struct {
	name string
	story string
	details Details
}

var byteString = []byte(`{"name": "Steve", "story": "This is the details"}`)

func CheckJsonResponse() {

	// var myInterface Person
	// json.Unmarshal(byteString, &myInterface)
	// fmt.Printf("%+v", myInterface)

	unknownStruct := make(map[string]interface{})

	json.Unmarshal(byteString, &unknownStruct)
	// fmt.Printf("%+v", unknownStruct)

	for key, value := range unknownStruct {
		fmt.Printf("%v : %v\n", key, value)
	}
}

