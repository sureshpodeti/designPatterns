package commands

import (
	"commandPattern/receivers"
)

type OffCommand struct {
	Device receivers.IDevice
}

func (off *OffCommand) Execute() {
	off.Device.Off()
}
