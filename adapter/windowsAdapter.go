package main

type WindowsAdapter struct {
	windows WindowsComputer
}

func (w *WindowsAdapter) InsertInSquarePort() {
	//Main logic
	//Taking Square and converting into circle
	w.windows.insertInCirclePort()
}
