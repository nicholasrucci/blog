package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type StatusResponse struct {
	StatusCode int
	Message    string
}

// SuccessStatus will send a 200 success status code
func SuccessStatus(w http.ResponseWriter, r *http.Request) error {
	res := StatusResponse{200, "The request was fulfilled."}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// ErrorStatus will send a 500 unexpected condition status code
func ErrorStatus(w http.ResponseWriter, r *http.Request) error {
	res := StatusResponse{500, "The server encountered an unexpected condition which prevented it from fulfilling the request."}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
