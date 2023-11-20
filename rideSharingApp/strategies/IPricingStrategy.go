package strategies

type IPricingStrategy interface {
	Execute(int, int) int
}
