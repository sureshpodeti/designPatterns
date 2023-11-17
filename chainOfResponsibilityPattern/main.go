package main

import "chainOfResponsibilityPattern/concrete"

func main() {
	cashier := &concrete.Cashier{}

	medicine := &concrete.Medicine{}

	medicine.SetNext(cashier)

	doctor := &concrete.Doctor{}
	doctor.SetNext(medicine)

	reception := &concrete.Reception{}
	reception.SetNext(doctor)

	patient := &concrete.Patient{}

	reception.Execute(patient)
}
