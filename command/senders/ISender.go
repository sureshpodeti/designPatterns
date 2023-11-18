package senders

type ISender interface {
	SetCommand()
	Execute()
}
