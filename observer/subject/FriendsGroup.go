package subject

import (
	"fmt"
	"observer/models"
	"observer/observer"
)

type FriendsGroup struct {
	Friends []*observer.Member
	Update  *models.Post
}

func (f *FriendsGroup) Join(m *observer.Member) {
	f.Friends = append(f.Friends, m)
}

func (f *FriendsGroup) Exit(m *observer.Member) {

	index := 0

	for i, fr := range f.Friends {
		if fr.ID == m.ID {
			index = i
			break
		}
	}
	f.Friends = append(f.Friends[:index], f.Friends[index+1:]...)
	fmt.Printf("Member left, ID: %d, name: %s\n", m.ID, m.Name)
}

func (f *FriendsGroup) Post(p *models.Post) {
	f.Update = p
	//Notify members
	f.Notify()
}
func (f *FriendsGroup) Notify() {
	for _, p := range f.Friends {
		p.GetUpdate(*f.Update)
	}
}
