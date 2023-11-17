package main

type TomatoTopping struct {
	Pizza IPizza
}

func (t *TomatoTopping) GetPrice() int {
	return t.Pizza.GetPrice() + 30
}
