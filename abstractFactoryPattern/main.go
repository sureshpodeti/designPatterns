package main

import "abstractFactoryPattern/abstract"

func main() {
	// Create a modern furniture set
	modernFactory := abstract.ModernFurnitureFactory{}
	modernChair := modernFactory.CreateChair()
	modernSofa := modernFactory.CreateSofa()
	modernCoffeeTable := modernFactory.CreateTable()

	modernChair.SitOn()
	modernSofa.LieOn()
	modernCoffeeTable.PutCoffee()

	// Create a victorian furniture set
	victorianFactory := abstract.VictorianFurnitureFactory{}
	victorianChair := victorianFactory.CreateChair()
	victorianSofa := victorianFactory.CreateSofa()
	victorianCoffeeTable := victorianFactory.CreateTable()

	victorianChair.SitOn()
	victorianSofa.LieOn()
	victorianCoffeeTable.PutCoffee()
}
