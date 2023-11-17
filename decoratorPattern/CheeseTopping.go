package main

type CheeseTopping struct {
	Pizza IPizza
}

func (c *CheeseTopping) GetPrice() int {
	return c.Pizza.GetPrice() + 25
}
