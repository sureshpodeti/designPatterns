package main

import (
	"chainOfResponsibility/Lift"
	"fmt"
)

func main() {

	/*
	  Chain of responsibility pattern (Chain of responsibility each will either process or pass it over to next)
	  1. Approval workflows
	  2. Customer support systems
	  3. Workflow systems (patient visiting doctor in hospital)
	  4. Lift ; need to go to source to destination; it will either stop at the intermediate floor, or go to next floor
	*/

	//patientStatus := &Hospital.PatientStatus{
	//	RegistrationDone:  false,
	//	DoctorCheckupDone: false,
	//	PaymentDone:       false,
	//}
	//patient := &Hospital.Patient{
	//	Name:          "Suresh Podeti",
	//	Age:           32,
	//	Complaint:     "Mild fever from last 2 days; No cold; No cough",
	//	PatientStatus: patientStatus,
	//}
	//
	//reception := &Hospital.Reception{}
	//doctor := &Hospital.Doctor{}
	//cashier := &Hospital.Cashier{}
	//
	//reception.SendTo(doctor)
	//doctor.SendTo(cashier)
	//
	//reception.Process(patient)
	//
	//fmt.Println("------------------------------------")
	//fmt.Println("Trying to process same patient again!")
	//fmt.Println()
	//reception.Process(patient)

	floor1 := &Lift.Floor{Number: 1}
	floor2 := &Lift.Floor{Number: 2}
	floor3 := &Lift.Floor{Number: 3}
	floor4 := &Lift.Floor{Number: 4}

	floor1.Goto(floor2)
	floor2.Goto(floor3)
	floor3.Goto(floor4)

	request := &Lift.Request{FloorNo: 3}

	floor1.Stop(request)

	fmt.Println("-------")

	request.FloorNo = 2

	floor1.Stop(request)
}
