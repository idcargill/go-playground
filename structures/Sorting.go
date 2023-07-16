package structures

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age int
}


// Sorting Interface 
type Family []Person

func(a Family) Len() int {
	return len(a)
}

func (a Family) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a Family) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}



func Sorting() {

	s := []int{
		100,
		5,
		1,
		50,
		75,
		15,
		40,
		60,
		20,
		75,
	}

	sort.Ints(s)
	fmt.Println(s)

	// Strings

	names := []string{
		"Frank",
		"zebra",
		"Zebra",
		"alice",
		"Alice",
	}
	sort.Strings(names)
	fmt.Println(names)


	// Descending: custom less method
	sort.Slice(s, func(i, j int) bool {
		return s[j] < s[i]
	})
	fmt.Println(s)

	// Sort structs
	family := []Person{
		{name: "Alice", age: 50},
		{name: "bob", age: 12},
		{name: "Ian", age: 42},
	}

	sort.Sort(Family(family))
	fmt.Println(family)
}