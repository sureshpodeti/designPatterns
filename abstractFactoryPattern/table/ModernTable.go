package table

import "fmt"

type ModernTable struct{}

func (m *ModernTable) PutCoffee() {
	fmt.Println("Putting coffee on modern table")
}
