package main

func main() {
	/*
		  We have two computer types;
			Apple {mac}, windows {ROG}
		 Mac latop has square port for charging
		 Windows has round port for charging
		Question?
			How can we charge Windows machine (ROG) using apple's square port
	*/
	mac := &Mac{}
	mac.InsertInSquarePort()

	windows := &ROG{}
	windows.insertInCirclePort()

	windowsAdapter := &WindowsAdapter{
		windows: windows,
	}

	windowsAdapter.InsertInSquarePort()
}
