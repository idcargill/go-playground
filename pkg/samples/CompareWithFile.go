package samples

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func CompareWithFile() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))	
	checkErr(err)

	// /home/ian/unicorn/go-play/data/mockJsonRes.json
	workingDirPath := "../../../../home/ian/unicorn/go-play"
	path1 := filepath.Join(dir, workingDirPath, "/data/mockJsonRes.json")
	path2 := filepath.Join(dir, workingDirPath, "/data/mockJsonRes2.json")
	
	// OPEN FILE
	f1 ,err := os.Open(path1)
	checkErr(err)
	defer f1.Close()

	f2, err := os.Open(path2)
	checkErr(err)
	defer f2.Close()

	// Read byte array
	result1, err := os.ReadFile(path1)
	result2, err := os.ReadFile(path2)
	checkErr(err)
	
	// Unmashall to JSON

	// unknown struct
	parsed1 := make(map[string]interface{})
	parsed2 := make(map[string]interface{})
	
	json.Unmarshal(result1, &parsed1)
	json.Unmarshal(result2, &parsed2)
	
	// COMPARE
	for key, value := range parsed1 {
		r2Value, ok := parsed2[key]
		// Key Missing
		if !ok {
			fmt.Printf("Test Failed: \"%v\" is missing but expected\n", key)
			continue
		}
		// Value not matching
		if value != r2Value {
			fmt.Printf("Test Failed:\n Expected: %v: %v\n Received: %v: %v\n", key,value, key, r2Value)
		} else {
			// Test passed
			fmt.Print(".")
		}
	}
	fmt.Println("DONE")
}