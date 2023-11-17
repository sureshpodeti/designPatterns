package publisher

import "observerPattern/subscriber"

type IStore interface {
	Register(observer subscriber.Observer)
	DeRegister(observer subscriber.Observer)
	NotifyAll()
}
