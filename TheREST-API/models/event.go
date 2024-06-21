package models

import "time"

// created a custom type that shape
type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
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

// get all avaliable events
func GetAllEvents() []Event {
	return events
}
