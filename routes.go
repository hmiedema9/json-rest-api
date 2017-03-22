package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

/* Route - type for URL structuring */
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

/* Routes - Array for all routes available */
type Routes []Route

/* New Router - Function for main.go */
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"LanguageIndex",
		"GET",
		"/languages",
		LanguageIndex,
	},
	Route{
		"LanguageShow",
		"GET",
		"/languages/{todoId}",
		LanguageShow,
	},
}
