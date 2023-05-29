package slices

import (
	"fmt"
	"sort"
)

type People struct {
	name string
	age int
}

type integers []int

var numbers []int
var objects []People


type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] > s[j]
}

// Sort objects in Slice
type sortedPeople []People

func (p sortedPeople) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p sortedPeople) Less(i, j int) bool {
	return p[i].age < p[j].age
}


func Sorting() {
	fmt.Println("===== SORTING =====")
	
	numbers = []int{5, 4,3,2,1}
	fmt.Println(numbers)
	
	// Sort Integer slice
	sort.Ints(numbers)
	fmt.Println(numbers)

	// Descending Int Slice
	nums := IntSlice{5, 10, 2, 400, 1, 8}
	sort.Sort(nums)
	fmt.Println(nums)


	// Sort Slice of Objects
	frank := People{name: "Frank", age: 10}
	jo := People{name: "JO", age: 15}
	puppy := People{name: "Puppy", age: 7}
	turtle := People{name: "Leonardo", age: 120}

	people := sortedPeople{frank, jo, puppy, turtle}
	sort.Slice(people, func (i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println(people)
}