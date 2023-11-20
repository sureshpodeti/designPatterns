package main

import (
	"awesomeProject9/models"
	"awesomeProject9/service"
	"fmt"
)

func main() {

	rider1 := &service.Rider{
		ID:     1,
		Person: models.Person{Name: "rider1"},
	}
	rider2 := &service.Rider{
		ID:     2,
		Person: models.Person{Name: "rider1"},
	}

	// Test case #1 origin: 50, destination: 60, noOfSeats: 1
	rider1.BookRide(1, 50, 60, 1)

	// Test case #1 origin: 50, destination: 60, noOfSeats: 2
	rider2.BookRide(2, 50, 60, 2)

	rider1.CloseRide()

	rider2.CloseRide()

	fmt.Println()
}
