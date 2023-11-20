package states

import (
	"awesomeProject9/service"
	"fmt"
)

type Started struct {
	ride *service.Ride
}

func (s *Started) StartRide() {
	fmt.Errorf("you can not start already started trip")
}

func (s *Started) CancelRide() {
	s.ride.Cancel()
	fmt.Println("Trip has been cancelled")
}

func (s *Started) EndRide() {
	s.ride.End()
	fmt.Println("Trip has been ended")

}
