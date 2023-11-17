package main

import (
	"observerPattern/publisher"
	"observerPattern/subscriber"
)

func main() {
	//Create producer
	store := publisher.NewStore()

	//Create subscriber
	customer := subscriber.NewCustomer(1)

	//register customer for item qty updates
	store.Register(customer)

	//store.RemoveItem()

	store.AddItem()
}
