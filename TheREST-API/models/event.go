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
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query) // stored in memory and reused
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(evt.Name, evt.Description, evt.Location, evt.DateTime, evt.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // to get an id that was inserted
	evt.ID = id
	return err

	// I'm not longer managing events as a local slice
	// appending a new event to events
	// events = append(events, evt)
}

// get all avaliable events
func GetAllEvents() ([]Event, error) {
	// to fetch an inserted events from the database
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) // we want to back a bunch of rows
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
