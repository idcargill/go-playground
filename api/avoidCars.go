package api

import (
	"fmt"
)

type CarBrand struct {
	name string
	models map[string]Model
}

type Model struct {
	name string
	avoid []int
}

var cars = make(map[string]CarBrand)

func ShouldCarBeAvoided(carMake string, model string, year int) bool {
	manufacturer := cars[carMake]
	badYears := manufacturer.models[model].avoid

	for _, v := range badYears {
		if year == v {
			return true
		}
	}
	return false
}

func GetCars() map[string]CarBrand {
	return cars
}

func GetAllBadYearsByModel(maker string, model string) []int {
	return cars[maker].models[model].avoid
}

func appendBadYears(years []int, year int) []int {
	years = append(years, year)
	return years
}

func CarModelExists(manufacturer string, model string) error {
	val, ok := cars[manufacturer]
	if !ok {
		return fmt.Errorf("Manufacturer not entered")
	} else {
		_, ok := val.models[model]
		if !ok {
			return fmt.Errorf("Model not found")
		}
	}
	return nil
}

 
func AddCarInfo(carMake string, model string, year int) {
		err := CarModelExists(carMake, model)
		if err != nil {
			badYears := []int{year}
			carModel := Model{avoid: badYears, name: model}
			carModels := make(map[string]Model)
			carModels[model] = carModel
		
			carBrand := CarBrand{name: carMake, models: carModels}
			cars[carMake] = carBrand
		}
}

func AddBadYear(manufacturer string, model string, year int) {
	car := cars[manufacturer].models[model]
	car.avoid = appendBadYears(car.avoid, year)
	cars[manufacturer].models[model] = car
	fmt.Printf("%+v", cars)
}

