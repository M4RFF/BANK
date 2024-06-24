package models

import (
	"theRestApi/db"
	"time"
)

// created a custom type that shape
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

// create a method that allows us
// to interact with our type's data

var events = []Event{}

func (evt Event) Save() error {
	// LATER: add it to a database
	query := `
	INSERT INTO events(name, secription, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(evt.Name, evt.Description, evt.Location, evt.DateTime, evt.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	evt.ID = id
	return err

	// appending a new event to events
	// events = append(events, evt)
}

// get all avaliable events
func GetAllEvents() []Event {
	return events
}
