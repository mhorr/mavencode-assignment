package main

import (
	"net/http"
)

// Route encapsulates a route for the webserver
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of Route objects
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	// Route{
	// 	"TodoIndex",
	// 	"GET",
	// 	"/todos",
	// 	TodoIndex,
	// },
	// Route{
	// 	"TodoShow",
	// 	"GET",
	// 	"/todos/{todoId}",
	// 	TodoShow,
	// },
	// Route{
	// 	"TodoCreate",
	// 	"POST",
	// 	"/todos",
	// 	TodoCreate,
	// },
	Route{
		"PersonCreate",
		"POST",
		"/person",
		PersonCreate,
	},
	Route{
		"PersonGet",
		"GET",
		"/person/{fullname}",
		PersonList,
	},
	Route{
		"PersonsQuery",
		"GET",
		"/persons/",
		PersonsQuery,
	},
	Route{
		"PersonsQuery",
		"GET",
		"/persons/{range}",
		PersonsQuery,
	},
}
