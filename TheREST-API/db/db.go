package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // a global variable

// initializing a data base
func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db") // open the connection

	if err != nil {
		panic("couldn't connect to database") // crash the app
	}

	DB.SetMaxOpenConns(10) // how many connections can be open
	DB.SetMaxIdleConns(5)  // how many connections we wanna keep open if nobody is using them

	CreateTables()
}

// it should create a bunch of tables in this db in case
// if they don't exist yet
func CreateTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS event(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL, 
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER 
	)
	`
	// to execute query statement
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("couldn't create events table")
	}

}
