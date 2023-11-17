package strategies

type SubtractStrategy struct{}

func (s *SubtractStrategy) Execute(a, b int) int {
	return a + b
}
