package structures

type Warrior struct {
	playerType string
	weapon Weapon
}

type Mage struct {
	playerType string 
	magic string
}


func GetHero(heroType string) string {

	if heroType != "warrior" && heroType != "mage" {
		return "fail"
	}

	switch heroType {
	case "warrior":
		return Warrior{
			playerType: "warrior",
			weapon: sword,
		}.playerType
	case "mage":
		return Mage{
			playerType: "mage",
			magic: "I do magic",
		}.playerType
	}

	return ""


}