package Hospital

import (
	"fmt"
)

type Doctor struct {
	Next Department
}

func (d *Doctor) Process(p *Patient) {
	if p.PatientStatus.DoctorCheckupDone {
		fmt.Println("Doctor: Checkup done already!")
		d.Next.Process(p)
		return
	}
	fmt.Printf("Doctor: Checking patient: %s\n", p.Name)
	p.PatientStatus.DoctorCheckupDone = true
	fmt.Println("Doctor: Checkup done!")
	fmt.Println()
	d.Next.Process(p)
}

func (d *Doctor) SendTo(dpt Department) {
	d.Next = dpt
}
