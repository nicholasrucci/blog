package main

import (
	"log"
	"net/http"
)

func main() {

	dbConnection()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
