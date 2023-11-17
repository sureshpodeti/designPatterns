package strategies

type DebitCard struct {
	Extra int
}

func (d *DebitCard) Pay(amount int) int {
	return amount + d.Extra
}
