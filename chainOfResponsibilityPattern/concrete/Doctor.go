package concrete

import (
	"fmt"
)

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckupDone {
		fmt.Println("Doctor checkup done!")
		d.next.Execute(p)
		return
	}

	fmt.Println("Doctor checking the patient...")
	p.DoctorCheckupDone = true
	d.next.Execute(p)
}

func (doc *Doctor) SetNext(d Department) {
	doc.next = d
}
