package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

func dbConnection() *sql.DB {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/blog?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
