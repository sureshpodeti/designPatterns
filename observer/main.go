package main

import (
	"observer/models"
	"observer/observer"
	"observer/subject"
)

func main() {

	/*
	 1. Newspaper/Article subscription
	 2. Group chat (Notify all except the person who posted)
	 3. Social media followers
	 4. YouTube subscription
	 3. Stock available notifications
	 4. Restaurant order status in food ordering apps {notified on order status change}
	*/

	group := &subject.FriendsGroup{}

	member1 := &observer.Member{
		ID:   1,
		Name: "Suresh Podeti",
	}

	member2 := &observer.Member{
		ID:   2,
		Name: "Sowjanya Podeti",
	}

	member3 := &observer.Member{
		ID:   3,
		Name: "Podeti Vasantha",
	}

	group.Join(member1)
	group.Join(member2)
	group.Join(member3)

	post1 := &models.Post{
		Text: "Let's meet tomorrow for my fair well!",
	}
	group.Post(post1)
	group.Exit(member3)

	post2 := &models.Post{
		Text: "Sorry, event got postponed!",
	}

	group.Post(post2)

}
