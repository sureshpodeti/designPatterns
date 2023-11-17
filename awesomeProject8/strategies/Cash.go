package strategies

type Cash struct {
	Extra int
}

func (c *Cash) Pay(amount int) int {
	return amount + c.Extra
}
