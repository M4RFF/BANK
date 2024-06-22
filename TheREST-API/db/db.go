package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db") // open the connection

	if err != nil {
		panic("couldn't connect to database")
	}

	DB.SetMaxOpenConns(10) // we have a pool of outgoing connections
	DB.SetMaxIdleConns(5)  // how many connections we wanna keep open if nobody is using them
}
