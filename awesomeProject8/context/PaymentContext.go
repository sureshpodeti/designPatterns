package context

import "awesomeProject8/strategies"

type PaymentContext struct {
	PaymentStrategy strategies.IPayStrategy
}

func (p *PaymentContext) SetPaymentStrategy(s strategies.IPayStrategy) {
	p.PaymentStrategy = s
}

func (p *PaymentContext) Pay(amount int) int {
	return p.PaymentStrategy.Pay(amount)
}
