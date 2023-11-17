package context

import "awesomeProject8/strategies"

type ArithmeticContext struct {
	Strategy strategies.IArithmeticStrategy
}

func (ar *ArithmeticContext) SetStrategy(strategy strategies.IArithmeticStrategy) {
	ar.Strategy = strategy
}

func (ar *ArithmeticContext) Execute(a, b int) int {
	return ar.Strategy.Execute(a, b)
}
