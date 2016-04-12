package main

import (
	"log"
	"net/http"
)

func main() {
	envVars()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
