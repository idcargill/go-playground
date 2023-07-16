package main

import (
	_ "example/structures"
	"fmt"
)

type PersonI interface {
	getName () string
}

type Person struct {
	name string
}

func (p *Person) getName() string {
	return p.name
}

func main() {
	// structures.Structures()

	jo := Person{name: "Jo"}
	
	name := jo.getName()

	fmt.Println(name)
}