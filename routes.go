package main

import (
	"net/http"
)

// Route is a struct contating the name,
// method, pattern, and handler function
// for the route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes holds an array of of type Route
type Routes []Route

var routes = Routes{
	Route{
		"AdminPostIndex",
		"GET",
		"/admin/posts",
		AdminPostIndex,
	},
	Route{
		"AdminPostNew",
		"GET",
		"/admin/posts/new",
		AdminPostNew,
	},
	Route{
		"AdminPostEdit",
		"GET",
		"/admin/posts/{[postID]}/edit",
		AdminPostEdit,
	},
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
		"/api/posts/{postID}",
		APIPostUpdate,
	},
	Route{
		"PostDelete",
		"DELETE",
		"/api/posts/{postID}",
		APIPostDelete,
	},
}
