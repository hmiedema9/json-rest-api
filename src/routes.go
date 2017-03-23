package main

import "net/http"

/* Route - type for URL structuring */
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

/* Routes - Array for all routes available */
type Routes []Route

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
		"/languages/{languageId}",
		LanguageShow,
	},
	Route{
		"LanguageCreate",
		"POST",
		"/languages",
		LanguageCreate,
	},
}
