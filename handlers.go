package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AdminPostIndex renders html to add a new post
func AdminPostIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/views/admin/posts.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
}

// AdminPostNew renders html to add a new post
func AdminPostNew(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/views/admin/newPost.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
}

// PostShow renders html file
func PostShow(w http.ResponseWriter, r *http.Request) {
	// body, _ := ioutil.ReadFile("./public/post.html")
	// fmt.Fprintf(w, string(body))
	var (
		id      int
		title   string
		content string
		posted  string
	)

	vars := mux.Vars(r)
	postID := vars["postID"]

	db := dbConnection()
	rows, err := db.Query("SELECT * FROM posts WHERE ID = ?", postID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &title, &content, &posted)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t, err := template.ParseFiles("public/views/post.html")
	if err != nil {
		log.Fatal(err)
	}
	data := &Post{Title: title, Content: content, Posted: posted}

	t.Execute(w, data)
}

// APIPostIndex queries the database for all of the posts,
// appends each row to an array of post, and then
// renders all of them as JSON
func APIPostIndex(w http.ResponseWriter, r *http.Request) {
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

// APIPostShow queries the database for a specific post by
// it's ID. It will then create a Post from the values
// that were grabbed and render the Post as JSON
func APIPostShow(w http.ResponseWriter, r *http.Request) {
	var (
		p       Post
		id      int
		title   string
		content string
		posted  string
	)

	vars := mux.Vars(r)
	postID := vars["postID"]

	db := dbConnection()
	rows, err := db.Query("SELECT * FROM posts WHERE ID = ?", postID)
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

// APIPostCreate accepts a JSON object that includes the
// attributes `title` and `content`. It will then insert
// a new row to the database containing that data. After
// adding the new row, PostCreate responds with a JSON
// object of the created Post
func APIPostCreate(w http.ResponseWriter, r *http.Request) {
	var (
		post    Post
		id      int
		title   string
		content string
		posted  string
	)

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

	db := dbConnection()
	_, err = db.Query("INSERT INTO posts (title, content) values (?, ?)", post.Title, post.Content)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM posts ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &title, &content, &posted)
		if err != nil {
			log.Fatal(err)
		}
		post = Post{id, title, content, posted}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	JSONHandler(w, r)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Fatal(err)
	}
}

// APIPostUpdate receives data from the client that will replace
// a post with the specified ID that is sent over as well.
func APIPostUpdate(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	postID := vars["postID"]

	db := dbConnection()
	_, err = db.Query("UPDATE posts SET title=?, content=? WHERE ID = ?", post.Title, post.Content, postID)
	if err != nil {
		log.Fatal(err)
		ErrorStatus(w, r)
	}
	defer db.Close()
	SuccessStatus(w, r)
}

// APIPostDelete queries the database for the ID specified
// by the client and deletes it.
// TODO: return `x` umong delete
func APIPostDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["postID"]

	db := dbConnection()
	_, err := db.Query("DELETE FROM posts WHERE ID = ?", postID)
	if err != nil {
		log.Fatal(err)
		ErrorStatus(w, r)
	}
	defer db.Close()
	SuccessStatus(w, r)

}
