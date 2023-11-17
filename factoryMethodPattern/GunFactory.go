package main

type GunFactory struct{}

func NewGun(typ string) IGun {
	switch typ {
	case "AK47":
		return &AK47{}
	case "Musket":
		return &Musket{}

	default:
		return nil
	}

	return nil
}
