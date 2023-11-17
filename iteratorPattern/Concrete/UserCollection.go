package Concrete

import "iteratorPattern/Interfaces"

type UserCollection struct {
	Users []*User
}

func (uc *UserCollection) CreateIterator() Interfaces.Iterator {
	return &UserIterator{Users: uc.Users, next: -1}
}
