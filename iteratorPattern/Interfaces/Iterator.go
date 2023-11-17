package Interfaces

import "iteratorPattern/Concrete"

type Iterator interface {
	Next() *Concrete.User
	Prev() *Concrete.User
}
