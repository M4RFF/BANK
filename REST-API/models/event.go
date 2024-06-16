package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (evt Event) Save() { // save an event
	// later: add it to a database
	events = append(events, evt) // added a new event (evt)
}

func GetAllEvents() []Event { // not for existing events, it is for availible events
	return events
}
