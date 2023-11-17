package senders

import "commandPattern/commands"

type TVButton struct {
	Command commands.ICommands
}

func (t *TVButton) Click() {
	t.Command.Execute()
}
