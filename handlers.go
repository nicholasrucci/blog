package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var posts Posts

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func PostIndex(w http.ResponseWriter, r *http.Request) {
	JSONHandler(w, r)
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Fatal(err)
	}
}

func PostShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	fmt.Fprintln(w, "Post show:", postId)
}

func PostCreate(w http.ResponseWriter, r *http.Request) {

}
