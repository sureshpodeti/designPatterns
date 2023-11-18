package Hospital

type PatientStatus struct {
	RegistrationDone  bool
	DoctorCheckupDone bool
	PaymentDone       bool
}

type Patient struct {
	Name          string
	Age           int
	Height        int
	Weight        int
	Complaint     string
	PatientStatus *PatientStatus
}
