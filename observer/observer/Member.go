package observer

import (
	"fmt"
	"observer/models"
)

type Member struct {
	ID   int
	Name string
}

func (m *Member) GetUpdate(p models.Post) {
	fmt.Printf("Received an update by member %d, name %s, and message: %s\n", m.ID, m.Name, p.Text)
}
