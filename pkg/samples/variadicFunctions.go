package samples

import (
	"fmt"
)


func VariadicFuncExample(text ...string) {
	// Function takes any number of args

	for _, word := range text {
		switch word {
		case "kitten":
			fmt.Println("YOU GOT A KITTEN")
		default:
			fmt.Println(word)
		}
	}
}