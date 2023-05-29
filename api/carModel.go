package api

import (
	"fmt"
	"sort"
)

type CarModel struct {
	manufacturer string
	model string
	yearsToAvoid []int
}

func (c *CarModel) getManufacturer() string {
	return c.manufacturer
}

func (c *CarModel) getModel() string {
	return c.model
}

func (c *CarModel) getYearsToAvoid() []int {
	return c.yearsToAvoid
}

type ByYear []int

func (a ByYear) Less(i, j int) bool {
	return a[i] < a[j]
}

func (c *CarModel) addBadYear(year int) {
	years := append(c.getYearsToAvoid(), year)
	sort.Slice(years, func(i, j int) bool {
		return years[i] < years[j]	
	})
	c.yearsToAvoid = years
}

func CarBuilder(manufacturer string, model string) CarModel {
	emptyArr := make([]int, 0)
	return CarModel{manufacturer: manufacturer, model: model, yearsToAvoid: emptyArr}
}

func CarModelStuff() {
	corolla := CarBuilder("Toyota", "Corolla")

	corolla.addBadYear(2006)
	corolla.addBadYear(2005)
	fmt.Println("**", corolla.getYearsToAvoid())
}