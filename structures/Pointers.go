package structures

import "fmt"

type Weapon struct {
	name string
	damage int
}


type Hero struct {
	name string
	weapon Weapon
}

func (h *Hero) Attack() {
	fmt.Printf("I am %v and do %v damage with my %s\n", h.name, h.weapon.damage, h.weapon.name)
}

var sword = Weapon{
	name: "Sword of Doom",
	damage: 50,
}

var club = Weapon{
	name: "Little Stick",
	damage: 5,
}

func Pointers() {
	warrior := Hero{
		name: "Dragon Slayer",
		weapon: sword,
	}

	weakPlayer := Hero{
		name: "Clucky",
		weapon: club,
	}

	warrior.Attack()
	weakPlayer.Attack()

	warriorWeapon := &warrior.weapon
	*warriorWeapon = club
	

	warrior.Attack()

	nums := []int{5,10,100,5000}

	p := &nums[3]

	fmt.Println("Pointer location: ", p)
	fmt.Println("Pointer Value: ", *p)

	*p = 999
	nums[3] = 888
	fmt.Println(nums)

	fmt.Println("Pointer location: ", p)
	fmt.Println("Pointer Value: ", *p)

}