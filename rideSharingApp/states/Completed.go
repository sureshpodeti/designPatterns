package states

import (
	"awesomeProject9/service"
	"fmt"
)

type Completed struct {
	ride *service.Ride
}

func (c *Completed) StartRide() {
	fmt.Println("You can not start completed trip")

}

func (c *Completed) CancelRide() {
	fmt.Println("You can not cancel already completed trip")
}

func (c *Completed) EndRide() {
	fmt.Println("You can not end already ended trip")
}
