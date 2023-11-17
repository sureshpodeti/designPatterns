package main

import (
	"fmt"
	"iteratorPattern/Concrete"
)

func main() {
	users := []*Concrete.User{{Name: "Suresh", Age: 32}, {Name: "Sowjanya", Age: 26}}

	userCollection := Concrete.UserCollection{Users: users}

	iterator := userCollection.CreateIterator()

	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Prev())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())

}
