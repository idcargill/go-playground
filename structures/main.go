package structures

import (
	"fmt"
)

var msg = "======== Structures ======== "


func slices() {
	s := [100]int{}

	s[1] = 100
		a := s[:2]
	fmt.Println("Len A:", len(a), "cap A:", cap(a))

	// Make
	fmt.Println("-- Make Slice:")
	m := make([]int, 5)

	fmt.Println(len(m))

	// Nested slices
	TicTacToe()
}


func Structures() {
	fmt.Println(msg)
	slices()
	Maps()
	Sorting()
	Pointers()
}