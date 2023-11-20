package service

import (
	"awesomeProject9/enum"
)

var PricingPerKM = 20

type Ride struct {
	ID          int
	Origin      int
	Destination int
	NoOfSeats   int
	status      int
}

func (r *Ride) Start() {
	r.status = enum.STARTED
}

func (r *Ride) Cancel() {
	r.status = enum.CANCELLED
}

func (r *Ride) End() {
	r.status = enum.COMPLETED
}
