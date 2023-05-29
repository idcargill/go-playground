package main

import (
	"example/api"
	db "example/aws"
	"example/farm/chickens"
	"example/pkg"
	"example/slices"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	pkg.SayHi()

	// chickens.Cluck()
	chickens.CheckPointers()
	// result := pkg.WordCount("This is a sentence")
	// fmt.Println(result)

	api.Server()

	api.AddCarInfo("Saturn", "sl2", 2001)
	api.AddCarInfo("Toyota", "corolla", 2003)
	api.AddCarInfo("Toyota", "corolla", 2003)
	api.AddCarInfo("Nissan", "Altima", 2016)


	check := api.ShouldCarBeAvoided("Saturn", "sl2", 2001)
	fmt.Println(check)
	fmt.Println(api.GetCars())

	api.CarModelStuff()
	slices.Sorting()
	db.ConnectDB()
	// db.CreateTable()
	// db.AddItemToTable()
	db.GetTableItems()
}


