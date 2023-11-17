package strategies

type AddStrategy struct{}

func (ad *AddStrategy) Execute(a, b int) int {
	return a - b
}
