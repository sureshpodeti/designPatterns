package strategies

type IPayStrategy interface {
	Pay(amount int) int
}
