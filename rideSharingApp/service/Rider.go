package service

import (
	"awesomeProject9/context"
	"awesomeProject9/enum"
	"awesomeProject9/models"
	"awesomeProject9/strategies"
	"fmt"
)

var pricePerKm = 20

type Rider struct {
	Person      models.Person
	ID          int
	CurrentRide *Ride
}

func (r *Rider) BookRide(ID, origin, destination, noOfSeats int) bool {
	if origin >= destination {
		fmt.Println("wrong values for origin and destination, can't book ride!")
		return false
	}
	ride := &Ride{ID: ID,
		Origin:      origin,
		Destination: destination,
		NoOfSeats:   noOfSeats,
		status:      enum.STARTED,
	}
	r.CurrentRide = ride
	fmt.Printf("Ride: %v booked successful\n", ride.ID)
	return true
}

func (r *Rider) UpdateRide(origin, destination, noOfSeats int) {
	//Yet to be implemented
}
func (r *Rider) CloseRide() {
	if r.CurrentRide == nil {
		fmt.Println("There is no current side to close!")
		return
	}
	strategy1 := &strategies.Strategy1{NoOfSeats: r.CurrentRide.NoOfSeats}
	strategy2 := &strategies.Strategy2{}

	var paymentContext context.PaymentContext

	travelDistance := r.CurrentRide.Destination - r.CurrentRide.Origin

	if r.CurrentRide.NoOfSeats == 1 {
		paymentContext.SetStrategy(strategy2)
	} else {
		paymentContext.SetStrategy(strategy1)
	}
	fmt.Println(paymentContext.GetRideFare(travelDistance, pricePerKm))

	r.CurrentRide.End()
}
