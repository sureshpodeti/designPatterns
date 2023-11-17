package chair

import "fmt"

type VictorianChair struct{}

func (v *VictorianChair) SitOn() {
	fmt.Println("Sitting on victorian Chair")
}
