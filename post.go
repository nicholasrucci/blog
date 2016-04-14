package main

// Post is a struct including
// attributes that pertain to
// a single Post
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Posted  string `json:"posted"`
}

// Posts is a type that holds an
// array of type Post
type Posts []Post
