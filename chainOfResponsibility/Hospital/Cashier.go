package Hospital

import (
	"fmt"
)

type Cashier struct {
	Next Department
}

func (c *Cashier) Process(p *Patient) {
	if p.PatientStatus.PaymentDone {
		fmt.Println("Cashier: Payment already done!")
		return
	}
	fmt.Println("Cashier: Patient doing payment...")
	p.PatientStatus.PaymentDone = true
	fmt.Println("Cashier: Payment done!")
}

func (c *Cashier) SendTo(d Department) {
	c.Next = d
}
