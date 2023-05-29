package chickens

type Chicken struct {
	name string
	age int
}

func ChickenBuilder(name string, age int) Chicken {
	newChicken := Chicken{name, age}
	return newChicken
}
