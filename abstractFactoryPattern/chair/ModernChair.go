package chair

import "fmt"

type ModernChair struct {
}

func (m *ModernChair) SitOn() {
	fmt.Println("Sitting on modern chair")
}
