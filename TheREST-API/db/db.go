package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // a global variable

// initializinga data base
func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db") // open the connection

	if err != nil {
		panic("couldn't connect to database") // crash the app
	}

	DB.SetMaxOpenConns(10) // how many connections can be open
	DB.SetMaxIdleConns(5)  // how many connections we wanna keep open if nobody is using them
}
