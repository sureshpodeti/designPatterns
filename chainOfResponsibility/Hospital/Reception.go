package Hospital

import (
	"fmt"
)

type Reception struct {
	Next Department
}

func (r *Reception) Process(p *Patient) {
	//Measure weight and height
	fmt.Println("Reception: Measuring weight and height of the patient")

	p.Height = 160 // cm
	p.Weight = 100 // kg

	// do the registration
	if p.PatientStatus.RegistrationDone {
		fmt.Println("Reception: Registration done already!")
		r.Next.Process(p)
		return
	}

	fmt.Printf("Reception: Weight: %d, height: %d |  Receptionist doing registration of %s\n", p.Weight, p.Height, p.Name)
	p.PatientStatus.RegistrationDone = true
	fmt.Println("Reception: Registration done!")
	fmt.Println()
	r.Next.Process(p)
}

func (r *Reception) SendTo(d Department) {
	r.Next = d
}
