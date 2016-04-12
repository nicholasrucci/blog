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

func SuccessStatus(w http.ResponseWriter, r *http.Request) error {
	res := StatusResponse{200, "The request was fulfilled."}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func ErrorStatus(w http.ResponseWriter, r *http.Request) error {
	res := StatusResponse{500, "The server encountered an unexpected condition which prevented it from fulfilling the request."}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
