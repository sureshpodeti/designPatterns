package senders

import "command/commands"

type Remote struct {
	command commands.ICommand
}

func (r *Remote) Execute() {
	r.command.Press()
}

func (r *Remote) SetCommand(command commands.ICommand) {
	r.command = command
}
