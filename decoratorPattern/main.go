package main

import "fmt"

func main() {

	// Initialize the base pizza
	basePizza := VegMania{}

	fmt.Println("base pizza cost: ", basePizza.GetPrice())

	// Put tomoto topping
	tomatoTop := TomatoTopping{Pizza: &basePizza}

	fmt.Println("Pizza with TomatoTopping cost: ", tomatoTop.GetPrice())

	cheeseTop := CheeseTopping{Pizza: &tomatoTop}

	fmt.Println("Pizza with cheese topping cost: ", cheeseTop.GetPrice())

}
