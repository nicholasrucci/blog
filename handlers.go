package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func PostIndex(w http.ResponseWriter, r *http.Request) {
	var (
		posts   Posts
		id      int
		title   string
		content string
		posted  time.Time
	)

	db := dbConnection()
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &title, &content, &posted)
		if err != nil {
			log.Fatal(err)
		}
		p := Post{id, title, content, posted}
		posts = append(posts, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
