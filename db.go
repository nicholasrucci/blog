package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

func dbConnection() {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1)/hello")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("Ping successful")
	}

	defer db.Close()
}
