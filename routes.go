package main

import (
	"net/http"
)

type ResponseWriter interface {
	Header() http.Header
	Write([]byte) (int, error)
	WriteHeader(int)
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"PostShow",
		"GET",
		"/posts/{postID}",
		PostShow,
	},
	Route{
		"PostIndex",
		"GET",
		"/api/posts",
		APIPostIndex,
	},
	Route{
		"PostCreate",
		"POST",
		"/api/posts",
		APIPostCreate,
	},
	Route{
		"PostShow",
		"GET",
		"/api/posts/{postID}",
		APIPostShow,
	},
	Route{
		"PostUpdate",
		"PUT",
		"/api/posts/{postId}",
		APIPostUpdate,
	},
	Route{
		"PostDelete",
		"DELETE",
		"/api/posts/{postId}",
		APIPostDelete,
	},
}
