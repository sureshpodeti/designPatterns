package commands

import "commandPattern/receivers"

type OnCommand struct {
	Device receivers.IDevice
}

func (on *OnCommand) Execute() {
	on.Device.On()
}
