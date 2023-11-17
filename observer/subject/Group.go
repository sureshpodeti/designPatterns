package subject

import (
	"observer/models"
	"observer/observer"
)

type Group interface {
	Join(m *observer.Member)
	Exit(m *observer.Member)
	Post(p *models.Post)
	Notify()
}
