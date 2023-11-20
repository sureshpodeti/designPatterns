package context

import (
	"awesomeProject9/strategies"
)

type PaymentContext struct {
	PricingStrategy strategies.IPricingStrategy
}

func (p *PaymentContext) SetStrategy(pr strategies.IPricingStrategy) {
	p.PricingStrategy = pr
}

func (p *PaymentContext) GetRideFare(distance, pricePerKm int) int {
	return p.PricingStrategy.Execute(distance, pricePerKm)
}
