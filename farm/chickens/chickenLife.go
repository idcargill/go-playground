package chickens

import (
	"fmt"
)

func Cluck() {
	fmt.Println("Cluck cluck cluck!")
}


var item int
var pointer *int

func CheckPointers() {
	item = 10
	pointer = &item

	fmt.Println(pointer)
	fmt.Println(*pointer)

	ChickenStruct()
}


var chickenArr [2]Chicken
var chickenSlice []Chicken

var chickenMap map[string]Chicken

func MapOfChickens() {
	chickenMap = make(map[string]Chicken)
	
	chickenMap["First"] = ChickenBuilder("One", 1)
	chickenMap["Second"] = ChickenBuilder("second", 2)

	fmt.Printf("%v", chickenMap)

	fmt.Println("\nRange")
	for k, _ := range chickenMap {
		fmt.Println(k)
	}
}

func ChickenStruct() {
	Frank := ChickenBuilder("Frank", 3)
	Lilian := ChickenBuilder("Lilian", 10)
	Rodeo := ChickenBuilder("Rodeo", 4)

	chickenSlice = append(chickenSlice, Frank, Lilian, Rodeo)
	
	chickenArr[0] = Lilian
	chickenArr[1] = Frank
	
	fmt.Println("Iterate array")
	for _,v := range chickenArr {
		fmt.Printf("Chickens: %+v\n", v)
	}

	MapOfChickens()
}
