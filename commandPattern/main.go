package main

import (
	"commandPattern/commands"
	"commandPattern/receivers"
	"commandPattern/senders"
)

func main() {

	//Receiver: TV; defualt off status
	tv := &receivers.TV{}

	//Commands
	onCmd := &commands.OnCommand{Device: tv}
	offCmd := &commands.OffCommand{Device: tv}

	// Sender: TV Button
	tvBtnOn := &senders.TVButton{Command: onCmd}
	tvBtnOn.Click()

	tvBtnOff := &senders.TVButton{Command: offCmd}
	tvBtnOff.Click()
}
