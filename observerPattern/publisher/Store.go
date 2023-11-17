package publisher

import (
	"observerPattern/subscriber"
)

type Store struct {
	Subscribers []subscriber.Observer
	ItemQty     int
}

func NewStore() *Store {
	return &Store{ItemQty: 0}
}
func (s *Store) Register(observer subscriber.Observer) {
	s.Subscribers = append(s.Subscribers, observer)
}
func (s *Store) DeRegister(observer subscriber.Observer) {
	indx := 0

	for i := 0; i < len(s.Subscribers); i++ {
		if s.Subscribers[i].GetID() == observer.GetID() {
			indx = i
			break
		}
	}

	//de-register logic
	l := s.Subscribers[0:indx]
	r := s.Subscribers[indx+1:]

	var result []subscriber.Observer
	result = append(result, l...)
	result = append(result, r...)
	s.Subscribers = result
}

func (s *Store) AddItem() {
	s.ItemQty++

	if s.ItemQty == 1 {
		s.NotifyAll()
	}
}

func (s *Store) RemoveItem() {
	if s.ItemQty == 0 {
		panic("can't remove an item from empty stock")
	}
	s.ItemQty--
}
func (s *Store) NotifyAll() {
	for _, observer := range s.Subscribers {
		observer.Update(s.ItemQty)
	}

}
