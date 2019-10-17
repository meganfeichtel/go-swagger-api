package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	// swagger:operation GET / Index getIndex
	//
	// Get the home page of the todo lists
	//
	// ---
	// responses:
	//   200:
	//     description: Get the homepage of the todo lists
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	// swagger:operation GET /todo TodoIndex getTodoIndex
	//
	// Get the list of todo items stored in the database
	//
	// ---
	// responses:
	//   200:
	//     description: Get the list of todo items stored in the database
	Route{
		"TodoIndex",
		"GET",
		"/todo",
		TodoIndex,
	},

	// swagger:operation GET /todo/{todoId} TodoShow getTodoShow
	//
	// Get the specified todo item
	//
	// ---
	// responses:
	//   200:
	//     description: Get the specified todo item
	Route{
		"TodoShow",
		"GET",
		"/todo/{todoId}",
		TodoShow,
	},

	// swagger:operation POST /todo TodoCreate postTodoCreate
	//
	// Create a new todo item
	//
	// ---
	// responses:
	//   200:
	//     description: Create a new todo item
	Route{
		"TodoCreate",
		"POST",
		"/todo",
		TodoCreate,
	},
}
