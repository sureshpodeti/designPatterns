package strategies

type CreditCard struct {
	Extra int
}

func (c *CreditCard) Pay(amount int) int {
	return amount + c.Extra
}
