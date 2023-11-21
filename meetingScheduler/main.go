package main

import "fmt"

func main() {

	/*
		Goal(s) of the system
		■
		Book a room for the meeting and return the name of the room booked
		■
		History of the meetings booked
		Write an API which accepts a start time and an end time, and returns the
		meeting room name for the booked scheduled time
	*/

	room1 := NewRoom("Room1")

	scheduler1 := &Scheduler{room: room1}

	scheduler2 := &Scheduler{room: room1}

	scheduler1.ScheduleMeeting(1000, 1100)

	scheduler2.ScheduleMeeting(1100, 1115)

	meetings := room1.GetHistory()

	for i, m := range meetings {
		fmt.Printf("meeting #%d: start=%d, end=%d\n", (i + 1), m.start, m.end)
	}

}
