package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (evt Event) Save() {
	// later: add it to a database
	events = append(events, evt) // added a new event (evt)
}
