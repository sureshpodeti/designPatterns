package states

import (
	"awesomeProject9/service"
	"fmt"
)

type Idle struct {
	ride *service.Ride
}

func (idl *Idle) StartRide() {
	idl.ride.Start()
	fmt.Println("Trip has been started")
}

func (idl *Idle) CancelRide() {
	fmt.Println("You can not cancel trip that does not start")
}

func (idl *Idle) EndRide() {
	fmt.Println("You can not end trip that does not start")
}
