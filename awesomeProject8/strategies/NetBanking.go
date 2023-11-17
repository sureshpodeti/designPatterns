package strategies

type NetBanking struct {
	Extra int
}

func (n *NetBanking) Pay(amount int) int {
	return amount + n.Extra
}
