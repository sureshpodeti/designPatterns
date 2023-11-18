package commands

import "command/receivers"

type OnCommand struct {
	Device receivers.IDevice
}

func (o *OnCommand) Press() {
	o.Device.On()
}
