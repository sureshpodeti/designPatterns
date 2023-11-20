package strategies

type Strategy1 struct {
	NoOfSeats int
}

func (s *Strategy1) Execute(km, amountPerKm int) int {
	value := float64(km) * float64(amountPerKm) * (0.75) * float64(s.NoOfSeats)
	return int(value)
}
