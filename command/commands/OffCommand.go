package commands

import "command/receivers"

type OffCommand struct {
	Device receivers.IDevice
}

func (o *OffCommand) Press() {
	o.Device.Off()
}
