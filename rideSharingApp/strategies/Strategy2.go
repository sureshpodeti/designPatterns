package strategies

type Strategy2 struct{}

func (s *Strategy2) Execute(km, amountPerKm int) int {
	return km * amountPerKm
}
