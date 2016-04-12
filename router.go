package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter is creating the router using that
// is being used throughout the project using
// gorilla mux
func NewRouter() *mux.Router {
	// create new router
	router := mux.NewRouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router

}

// JSONHandler is being put in all of the handler functions to
// make the Content Type of the responses as JSON.
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
