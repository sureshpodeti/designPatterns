package subscriber

import "fmt"

type Customer struct {
	ID int
}

func NewCustomer(id int) *Customer {
	return &Customer{id}
}
func (c *Customer) Update(qty int) {
	fmt.Printf("Customer Received update on item qty: %d\n", qty)
}

func (c *Customer) GetID() int {
	return c.ID
}
