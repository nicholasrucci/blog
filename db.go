package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

func dbConnection() {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
