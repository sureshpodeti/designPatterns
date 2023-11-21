package main

type Scheduler struct {
	room *Room
}

func (s *Scheduler) ScheduleMeeting(startTime, endTime int) bool {
	return s.room.BookRoom(startTime, endTime)
}
