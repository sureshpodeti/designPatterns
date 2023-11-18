package Lift

type IFloor interface {
	Stop(r *Request)
	Goto(fl *Floor)
}
