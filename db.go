package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

// dbConnection opens a new database connection and
// returns a pointer to the sql.DB
func dbConnection() *sql.DB {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
