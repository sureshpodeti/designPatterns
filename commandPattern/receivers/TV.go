package receivers

import (
	"fmt"
)

type TV struct {
	IsOn bool
}

func (t *TV) On() {
	t.IsOn = true
	fmt.Println("TV on")
}

func (t *TV) Off() {
	t.IsOn = false
	fmt.Println("TV off")
}
