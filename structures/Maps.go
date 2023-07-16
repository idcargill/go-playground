package structures

import (
	"fmt"
)

func Maps() {
	fmt.Println("-- Maps --")
	var mString map[string]string

	mInt := make(map[int]string)
	mString = make(map[string]string)
	
	mInt[1] = "shark 1"
	mString["Kitten"] = "Tiger"

	fmt.Printf("%v+ \n", mInt)
	fmt.Printf("%v+ \n", mString)
}