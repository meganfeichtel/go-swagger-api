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

	// swagger:operation GET / root getIndex
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

	// swagger:operation GET /todo todo getTodoIndex
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

	// swagger:operation GET /todo/{todoId} todo getTodoShow
	// ---
	// summary: List the specified todo item.
	// description: List the specified todo item
	// parameters:
	// - name: todoId
	//   in: path
	//   description: id of todo list item
	//   type: string
	//   required: true
	// responses:
	//   200:
	//     description: Get the specified todo item

	Route{
		"TodoShow",
		"GET",
		"/todo/{todoId}",
		TodoShow,
	},

	// swagger:operation POST /todo todo postTodoCreate
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
