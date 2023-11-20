package states

type IStatus interface {
	StartRide()
	CancelRide()
	EndRide()
}
