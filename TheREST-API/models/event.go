package models

import "time"

// created a custom type that shape
type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

// create a method that allows us
// to interact with our type's data

var events = []Event{}

func (evt Event) Save() {
	// LATER: add it to a database

	// appending a new event to events
	events = append(events, evt)
}
