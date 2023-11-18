package main

import (
	"fmt"
	"log"
)

func main() {

	/*
	  1. Document publish process  {Draft, Moderation, published}
	  2. Traffic lights { Green, Yellow, Red }
	  3. Food ordered
	*/

	vendingMachine := NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Enter money")
	var amount int
	fmt.Scanln(&amount)

	err = vendingMachine.InsertMoney(amount)

	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()

	if err != nil {
		log.Fatalf(err.Error())
	}

}
