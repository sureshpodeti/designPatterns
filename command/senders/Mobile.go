package senders

import "command/commands"

type Mobile struct {
	command commands.ICommand
}

func (m *Mobile) Execute() {
	m.command.Press()
}

func (m *Mobile) SetCommand(command commands.ICommand) {
	m.command = command
}
