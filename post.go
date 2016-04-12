package main

import ()

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Posted  string `json:"posted"`
}

type Posts []Post
