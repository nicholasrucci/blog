package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

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
		posted  string
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
	var (
		p       Post
		id      int
		title   string
		content string
		posted  string
	)

	vars := mux.Vars(r)
	postId := vars["postId"]

	db := dbConnection()
	rows, err := db.Query("SELECT * FROM posts WHERE ID = ?", postId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &title, &content, &posted)
		if err != nil {
			log.Fatal(err)
		}
		p = Post{id, title, content, posted}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	JSONHandler(w, r)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Fatal(err)
	}
}

func PostCreate(w http.ResponseWriter, r *http.Request) {
	var post Post

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &post); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
	}
	log.Print(post)
}
