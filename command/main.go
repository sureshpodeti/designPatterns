package main

import (
	"command/commands"
	"command/receivers"
	"command/senders"
)

func main() {

	/*
			1. Undo/Redo actions; Text editor
		    2. TV on/off
		    3. Game actions
	*/

	//senders; mobile, remote
	mobile := &senders.Mobile{}
	remote := &senders.Remote{}

	//receiver: tv
	tv := &receivers.Tv{}

	//commands
	onCommand := &commands.OnCommand{Device: tv}
	offCommand := &commands.OffCommand{Device: tv}

	mobile.SetCommand(onCommand)
	mobile.Execute()

	mobile.SetCommand(offCommand)
	mobile.Execute()

	mobile.SetCommand(offCommand)
	mobile.Execute()

	mobile.SetCommand(onCommand)
	mobile.Execute()

	mobile.SetCommand(onCommand)
	mobile.Execute()

	remote.SetCommand(onCommand)
	remote.Execute()

	remote.SetCommand(onCommand)
	remote.Execute()

	remote.SetCommand(offCommand)
	remote.Execute()
}
