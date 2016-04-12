package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
	"os"
)

// dbConnection opens a new database connection and
// returns a pointer to the sql.DB
func dbConnection() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
