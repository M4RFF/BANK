package models

import (
	"rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (evt Event) Save() error { // save an event
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
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
	// later: add it to a database
	// events = append(events, evt) // added a new event (evt)
}

func GetAllEvents() []Event { // not for existing events, it is for availible events

	return events
}
