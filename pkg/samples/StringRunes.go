package samples

import (
	"fmt"
)

// String is a slice of bytes encoded in UTF-8
// Rune is a character, an integer that represents Unicode point
// Single quotes are "Rune Literals"
//

func RunesLand() {
	s := "Hello"

	for _, letter := range s {
		fmt.Printf("Rune %x", letter)



	}
}