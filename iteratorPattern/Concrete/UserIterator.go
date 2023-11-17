package Concrete

type UserIterator struct {
	next  int
	Users []*User
}

func (uc *UserIterator) Next() *User {
	uc.next++
	if uc.next >= len(uc.Users) {
		panic("Can not increment")
	}
	user := uc.Users[uc.next]
	return user
}

func (uc *UserIterator) Prev() *User {
	uc.next--
	if uc.next < 0 {
		panic("Can not go left")
	}
	user := uc.Users[uc.next]
	return user
}
