package main

import (
	"fmt"
	"sort"
)

type Room struct {
	name     string
	meetings []*Meeting
}

func NewRoom(name string) *Room {
	return &Room{name: name}
}

func (r *Room) BookRoom(startTime, endTime int) bool {
	if startTime >= endTime {
		fmt.Println("Room can not be booked; as startTime >= endTime")
		return false
	}

	// Check for conflicts
	//First sort by startTime of the schedule
	sort.Slice(r.meetings, func(i, j int) bool {
		return r.meetings[i].start < r.meetings[j].start
	})

	for _, m := range r.meetings {
		if m.end > startTime || m.start > endTime {
			fmt.Println("Conflict: Room can not be booked, as there is a conflict!")
			return false
		}
	}

	meeting := &Meeting{start: startTime, end: endTime}

	r.meetings = append(r.meetings, meeting)

	fmt.Printf("Meeting booked, from: %d, to: %d\n", startTime, endTime)

	return true
}

func (r *Room) GetHistory() []*Meeting {
	return r.meetings
}
