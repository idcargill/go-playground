package main

import (
	. "example/pkg/samples"
	_ "example/structures"
)

type PersonI interface {
	getName() string
}

type Person struct {
	name string
}

func (p *Person) getName() string {
	return p.name
}

func main() {
	// structures.Structures()

	// jo := Person{name: "Jo"}

	// name := jo.getName()
	// fmt.Println(name)

	// api.CheckJsonResponse()

	// === Samples ===
	// VariadicFuncExample("fish", "kitten", "dog")
	// RunesLand()
	// RateLimiting()
	// Generics()
	CompareWithFile()
}
