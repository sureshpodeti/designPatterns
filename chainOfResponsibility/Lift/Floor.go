package Lift

import "fmt"

type Floor struct {
	Number int
	Next   IFloor
}

func (f *Floor) Stop(r *Request) {
	fmt.Printf("At floor: %d\n", f.Number)

	if f.Number == r.FloorNo {
		fmt.Printf("Stopped at requested floor number: %d\n", f.Number)
	} else {
		f.Next.Stop(r)
	}
}

func (f *Floor) Goto(fl *Floor) {
	f.Next = fl
}
