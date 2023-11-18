package receivers

import "fmt"

type Tv struct {
	IsOn bool
}

func (t *Tv) On() {
	if t.IsOn {
		fmt.Println("Tv: On already, ignoring the action")
		return
	}
	t.IsOn = true
	fmt.Println("Tv is switched on")
}

func (t *Tv) Off() {
	if !t.IsOn {
		fmt.Println("Tv: Off already, ignoring the action")
		return
	}
	t.IsOn = false
	fmt.Println("Tv is switched off")
}
