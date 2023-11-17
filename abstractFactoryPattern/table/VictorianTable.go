package table

import "fmt"

type VictorianTable struct{}

func (v *VictorianTable) PutCoffee() {
	fmt.Println("Putting cofee on victorian table")
}
