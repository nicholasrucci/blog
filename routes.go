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
		"PostIndex",
		"GET",
		"/api/posts",
		PostIndex,
	},
	Route{
		"PostCreate",
		"POST",
		"/api/posts",
		PostCreate,
	},
	Route{
		"PostShow",
		"GET",
		"/api/posts/{postId}",
		PostShow,
	},
	Route{
		"PostUpdate",
		"PUT",
		"/api/posts/{postId}",
		PostUpdate,
	},
	Route{
		"PostDelete",
		"DELETE",
		"/api/posts/{postId}",
		PostDelete,
	},
}
