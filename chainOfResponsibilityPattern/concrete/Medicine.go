package concrete

import (
	"fmt"
)

type Medicine struct {
	next Department
}

func (m *Medicine) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine done already")
		m.next.Execute(p)
		return
	}
	fmt.Println("Providing medicine to patient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Medicine) SetNext(d Department) {
	m.next = d
}
