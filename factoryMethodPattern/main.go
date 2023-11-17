package main

import "fmt"

func main() {
	factory := NewGun("AK47")

	fmt.Println(factory.GetName())

	factory1 := NewGun("Musket")

	fmt.Println(factory1.GetName())
}
