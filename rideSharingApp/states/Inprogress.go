package states

import (
	"awesomeProject9/service"
	"fmt"
)

type Inprogress struct {
	ride *service.Ride
}

func (i *Inprogress) StartRide() {
	fmt.Println("You can not cancel trip that does not start")
}

func (i *Inprogress) CancelRide() {
	i.ride.Cancel()
	fmt.Println("Trip has been cancelled!")
}

func (i *Inprogress) EndRide() {
	i.ride.End()
	fmt.Println("Trip has been ended")
}
