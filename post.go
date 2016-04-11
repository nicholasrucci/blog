package main

import (
	"time"
)

type Post struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Posted  time.Time `json:"posted"`
}

type Posts []Post
