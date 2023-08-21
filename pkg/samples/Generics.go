package samples

import (
	"encoding/json"
	"fmt"
)

func Generics() {

	// Unmarshall json data
	resJson := "{\"name\":\"Tiger\", \"size\": 2}"
	whateverStruct := make(map[string]interface{})
	var err error
	err = json.Unmarshal([]byte(resJson), &whateverStruct)
	if err != nil {
		fmt.Println(err)
	}

	for k, text := range whateverStruct {
		fmt.Printf("%s : %s type => %T\n", k, text, text)
	}

}

// Generics 
func MapKeys[k comparable, v any](m map[k]v) {
	for key, value := range m {
		fmt.Println(key, " : ", value)
	}
}