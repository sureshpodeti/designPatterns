package concrete

import (
	"fmt"
)

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment done already!")
		return
	}

	fmt.Println("Patient doing payment...")
	p.PaymentDone = true
}

func (c *Cashier) SetNext(d Department) {
	c.next = d
}
